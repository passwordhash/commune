import {defineStore} from "pinia";
import {computed, ref, watch} from "vue";
import axios from "axios";
import router from "@/router/index.js";

const LS_TOKEN = "token"
const LS_MY_ID = "accountId"

export const useUserStore = defineStore('user', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    const token = ref(localStorage.getItem(LS_TOKEN));
    watch(token, (newToken) => {
        localStorage.setItem(LS_TOKEN, newToken);
    })
    const accountId = ref(localStorage.getItem(LS_MY_ID));
    watch(accountId, (newId) => {
        localStorage.setItem(LS_MY_ID, newId);
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

    function setData(newToken, newMyId) {
        token.value = newToken
        accountId.value = newMyId
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
        localStorage.removeItem(LS_MY_ID)

        isAuthenticated.value = false
        router.push("/")
    }

    return {
        token,
        accountId,
        isAuthenticated,
        authorize,
        setData,
        register,
        logout
    }
})
