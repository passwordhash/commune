<script setup>
import {onKeyStroke, useWebSocket} from "@vueuse/core";
import {computed, onBeforeMount, onBeforeUnmount, onMounted, onRenderTracked, ref, warn, watch} from "vue";
import Message from "@/components/Message.vue";
import {useMessageStore} from "@/stores/message.js";
import {useUserStore} from "@/stores/user.js";
import router from "@/router/index.js";

let storeUser = useUserStore()
let storeMsg = useMessageStore()

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

    storeMsg.fetchMessages(storeUser.token)

    storeMsg.newMessage({
        text: inputMsg.value
    }, storeUser.token)
        .then(() => {
            send(inputMsg.value)
        })
        .then(() => {
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
    console.log(newVal)
    storeMsg.addMessage(
        JSON.parse(newVal)
    )
})
onBeforeMount(() => {
    let token = storeUser.token
    console.log("token: ", token)
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
    <div class="container container-chat my-4">
        <div class="chat-wrapper">
            <Message v-for="msg in storeMsg.messages"
                     :key="msg.id"
                     :msg="msg"
                     :is-owned="storeUser.accountId === msg.author.id"
                     :is-last="isLast(msg)"
                     @mounted="scrollToBottom"
            />

            <!--      <div class="message own-message">-->
            <!--        <div><small class="text-muted">You, 15:01</small></div>-->
            <!--        <div class="message-text">I'm fine, thanks!</div>-->
            <!--      </div>-->
            <div ref="bottom"></div>
        </div>
        <div class="input-group">
            <input
                v-model="inputMsg"
                @keydown.enter="submit"
                type="text"
                class="form-control"
                placeholder="Ваше сообщение ...">
            <div class="input-group-append">
                <button @click="submit" class="btn btn-primary" :disabled="!canSend">Отправить</button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.container-chat {
    padding-top: 25px;
}

.chat-wrapper {
    border: 1px solid #4CAF50;
    border-radius: 20px;
    background-color: #C8E6C9;
    height: 60vh;
    overflow: auto;
    padding: 20px;
    margin-bottom: 20px;
    display: flex;
    flex-direction: column;
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
    max-width: 400px;
}

.form-control:focus {
    border-color: #4CAF50;
    box-shadow: 0 0 0 0.2rem rgba(76, 175, 80, 0.25);
}

.input-group {
    max-width: 400px;
    float: right;
}

.form-control, .btn-primary {
    border-radius: 15px;
    height: 50px;
}
</style>