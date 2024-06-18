import {defineStore} from "pinia";
import {ref} from "vue";
import axios from "axios";

export const useUserStore = defineStore('user', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    const user = ref({})

    function authorize(user) {
        return axios.post(`${baseUrl}/api/sign-up`, user)
    }

    return {
        user,
        authorize
    }
})
