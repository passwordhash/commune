import {ref} from 'vue'
import {defineStore} from 'pinia'
import axios from "axios";

export const useMessageStore = defineStore('message', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    const messages = ref([])

    function fetchMessages(token) {
        return axios.get(`${baseUrl}/api/list`, {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
            .then(res => {
                messages.value = res.data
            })
    }

    function newMessage(message, token) {
        return axios.post(`${baseUrl}/api/new`, message, {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
            .catch(error => {
                console.log(error)
            })
    }

    function addMessage(message) {
        messages.value.push(message)
    }

    return {
        messages,
        fetchMessages,
        addMessage,
        newMessage
    }
})