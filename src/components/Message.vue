<script setup>
import {computed, onMounted} from "vue";
import moment from "moment";
import 'moment/src/locale/ru.js'


const emit = defineEmits(['mounted'])

const props = defineProps({
    isLast: {
        type: Boolean,
        default: false,
    },
    msg: {
        type: Object,
        required: true,
    },
    isOwned: {
        type: Boolean,
    }
})

onMounted(() => {
    moment.locale("ru");
    if (props.isLast) {
        emit('mounted')
    }
})

let formatMessageTime = computed (() => {
    // const now = moment();
    const messageTime = moment(props.msg.date);

    return messageTime.format('HH:mm');
})

</script>

<template>
    <div class="message-card"
         :class="{'own-message': props.isOwned}"
    >
        <div v-if="!isOwned" class="message-header">
            <span class="username">{{ msg.author.nickname }}</span>
        </div>
        <div class="message-content">
            <p class="wordwrap">{{ msg.text }}</p>
        </div>
        <div class="message-time">
            {{ formatMessageTime }}
        </div>
    </div>
</template>

<style scoped>

.own-message {
    align-self: end;
}

.message-card {
    background-color: #6c757d;
    color: white;
    padding: 10px;
    border-radius: 15px;
    min-width: 20%;
    max-width: 42%;
    margin: 10px 0;
    position: relative;
}

.message-card.own-message {
   background-color: #4CAF50;
    color: #B0B7BE;
}

.message-header {
    display: flex;
    align-items: center;
}

.username {
    font-size: 14px;
    color: #B0B7BE;
    font-weight: bold;
    margin-right: 5px;
}

.message-content {
    color: #f8f8f8 !important;
    font-size: 16px;
    margin-bottom: 3px;
}

.message-time {
    font-style: italic;
    position: absolute;
    bottom: 5px;
    right: 10px;
    font-size: 12px;
    color: #f8f8f8 !important;
}
</style>