import {ref, computed} from 'vue'
import {defineStore} from 'pinia'
import axios from "axios";

export const useMessageStore = defineStore('user', () => {
    const baseUrl = import.meta.env.VITE_API_BASE_DOMAIN

    const messages = ref([])

    function fetchMessages() {
        axios.get(`${baseUrl}/api/list`)
            .then(response => {
                messages.value = response.data
            })

            .catch(error => {
                console.log(error)
            })
    }

    function newMessage(message) {
        return axios.post(`${baseUrl}/api/new`, message)
            .catch(error => {
                console.log(error)
            })
    }

    function addMessage(message) {
        messages.value.push(message)
    }

    return {messages, fetchMessages, addMessage, newMessage}
})