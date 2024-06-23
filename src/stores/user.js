import {defineStore} from "pinia";
import {ref} from "vue";
import axios from "axios";

export const useUserStore = defineStore('user', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    const user = ref({})
    // const primaryColor = ref("")

    // increment
    // const doubleCount = computed(() => count.value * 2)
    function authorize(email, passcode) {
        return axios.post(`${baseUrl}/auth/sign-in`, {
            email: email,
            password: passcode
        })
    }

    function register(email, nickname) {
        return axios.post(`${baseUrl}/auth/sign-up`, {
            nickname: nickname,
            email: email
        })
    }

    return {
        user,
        authorize,
        register
    }
})
