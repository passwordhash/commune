import {defineStore} from "pinia";
import {computed, ref, watch} from "vue";
import axios from "axios";
import router from "@/router/index.js";

const LS_TOKEN = "token"

export const useUserStore = defineStore('user', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    const token = ref(localStorage.getItem(LS_TOKEN));
    watch(token, (newToken) => {
        localStorage.setItem(LS_TOKEN, newToken);
    })

    const isAuthenticated = ref(false)

    // increment
    // const doubleCount = computed(() => count.value * 2)
    function authorize(email, passcode) {
        return axios.post(`${baseUrl}/auth/sign-in`, {
            email: email,
            passcode: passcode
        })
    }

    function setToken(newToken) {
        token.value = newToken
        isAuthenticated.value = true
    }

    function register(email, nickname) {
        return axios.post(`${baseUrl}/auth/sign-up`, {
            nickname: nickname,
            email: email
        })
    }

    function logout() {
        localStorage.removeItem(LS_TOKEN)
        isAuthenticated.value = false
        router.push("/")
    }

    return {
        token,
        isAuthenticated,
        authorize,
        setToken,
        register,
        logout
    }
})
