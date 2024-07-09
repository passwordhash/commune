<script setup>
import {onKeyStroke, useWebSocket} from "@vueuse/core";
import {computed, onBeforeMount, onBeforeUnmount, onMounted, onRenderTracked, ref, warn, watch} from "vue";
import Message from "@/components/Message.vue";
import {useMessageStore} from "@/stores/message.js";
import {useUserStore} from "@/stores/user.js";
import router from "@/router/index.js";
import ChatHeader from "@/components/ChatHeader.vue";

let storeUser = useUserStore()
let storeMsg = useMessageStore()
let account = ref({})

const {status, data, send, open, close} = useWebSocket('ws://localhost:8090/ws/chat', {
    autoReconnect: {
        retries: 3,
        delay: 1000,
        onFailed() {
            alert('Failed to connect WebSocket after 3 retries')
        },
    },
})

const inputMsg = ref('')
const bottom = ref()
const isFirstMount = ref(true)

const submit = (e) => {
    e.preventDefault()

    if (!canSend.value) {
        return
    }

    // storeMsg.fetchMessages(storeUser.token)

    storeMsg.newMessage({
        text: inputMsg.value
    }, storeUser.getAccount().token)
        .then((res) => {
            console.log(JSON.stringify(res.data))
            console.log(res.data)
            send(JSON.stringify(res.data))
            inputMsg.value = ""
        })
}

const scrollToBottom = () => {
    let params = {}
    if (!isFirstMount.value) {
        params = {behavior: 'smooth'}
    }
    bottom.value.scrollIntoView(params)

    isFirstMount.value = false
}

const isLast = (msg) => {
    return msg.id === storeMsg.messages[storeMsg.messages.length - 1].id
}

const canSend = (computed(() => {
    return status.value === 'OPEN' && inputMsg.value !== ''
}))

watch(data, (newVal) => {
    storeMsg.addMessage(
        JSON.parse(newVal)
    )
})
onBeforeMount(() => {
    account.value = storeUser.getAccount()
    let token = account.value.token

    if (token === null || token === 0) {
        return router.push('/')
    }
    // TODO: проверка самого токена

    storeMsg.fetchMessages(token)
        .catch( err => {
            if (err.response.status === 401) {
                console.log("user unauthorized")
                storeUser.logout()
                router.push("/")
            }
        })
})
</script>

<template>
    <main>
        <div class="container container-chat mb-4">
            <ChatHeader
                :user="account"
            />
            <div class="chat-wrapper">
                <Message v-for="msg in storeMsg.messages"
                         :key="msg.id"
                         :msg="msg"
                         :is-owned="storeUser.getAccount().id === msg.author.id"
                         :is-last="isLast(msg)"
                         @mounted="scrollToBottom"
                />
                <div ref="bottom"></div>
            </div>
            <div class="input-group">
                <div class="input-group-prepend">
                    <button @click="submit" class="btn btn-primary" :disabled="!canSend">Отправить</button>
                </div>
                <input
                    v-model="inputMsg"
                    @keydown.enter="submit"
                    type="text"
                    class="form-control"
                    placeholder="Ваше сообщение ...">
            </div>
        </div>
    </main>
</template>

<style scoped>
.chat-wrapper {
    position: relative;
    z-index: 1000;
    border: 1px solid #8ccb8f;
    border-radius: 20px;
    background-color: #ddedde;
    height: 70vh;
    overflow: auto;
    padding: 20px;
    margin-bottom: 20px;
    display: flex;
    flex-direction: column;
    align-items: start;
}

.btn-primary {
    background-color: #4CAF50;
    border-color: #4CAF50;
}

.btn-primary:disabled {
    background-color: #8BC34A;
    border-color: #8BC34A;
}

.form-control {
    border-radius: 20px;
    border-color: #4CAF50;
    margin-bottom: 10px;
    height: 50px;
    width: 50%;
}

.form-control:focus {
    border-color: #4CAF50;
    box-shadow: 0 0 0 0.2rem rgba(76, 175, 80, 0.25);
}

.form-control, .btn-primary {
    border-radius: 15px;
    height: 50px;
}
</style>