import {defineStore} from "pinia";
import {computed, ref, watch} from "vue";
import axios from "axios";
import router from "@/router/index.js";

// const LS_TOKEN = "token"
// const LS_MY_ID = "accountId"
const LS_ACCOUNT = "userAccount"

export const useUserStore = defineStore('user', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    // const token = ref(localStorage.getItem(LS_TOKEN));
    // watch(token, (newToken) => {
    //     localStorage.setItem(LS_TOKEN, newToken);
    // })
    // const accountId = ref(localStorage.getItem(LS_MY_ID));
    // watch(accountId, (newId) => {
    //     localStorage.setItem(LS_MY_ID, newId);
    // })
    const account = ref(localStorage.getItem(LS_ACCOUNT))
    watch(account, (newData) => {
        localStorage.setItem(LS_ACCOUNT, (JSON.stringify(newData)))
    })

    const getAccount = () => {
        let accountStr = localStorage.getItem(LS_ACCOUNT)
        return JSON.parse(accountStr)
    }

    const isAuthenticated = ref(false)

    // increment
    // const doubleCount = computed(() => count.value * 2)
    function authorize(email, passcode) {
        return axios.post(`${baseUrl}/auth/sign-in`, {
            email: email,
            passcode: passcode
        })
    }

    function register(email, nickname) {
        return axios.post(`${baseUrl}/auth/sign-up`, {
            nickname: nickname,
            email: email
        })
    }

    function logout() {
        localStorage.removeItem(LS_ACCOUNT)

        isAuthenticated.value = false
        router.push("/")
    }

    function setData(newToken, newUserData) {
        account.value = newUserData
        account.value.token = newToken
        isAuthenticated.value = true
    }


    return {
        getAccount,
        isAuthenticated,
        authorize,
        setData,
        register,
        logout
    }
})
