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
        <div class="message-header">
            <span class="username">{{ msg.author.nickname }}</span>
        </div>
        <div class="message-content">
            <p>{{ msg.text }}</p>
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
    background-color: #2C2F33;
    color: white;
    padding: 10px;
    border-radius: 15px;
    width: 38%;
    margin: 10px 0;
    position: relative;
}

.message-header {
    display: flex;
    align-items: center;
}

.username {
    color: #D081C5;
    font-weight: bold;
    margin-right: 5px;
}

.message-content {
    margin: 5px 0;
}

.message-time {
    font-style: italic;
    position: absolute;
    bottom: 5px;
    right: 10px;
    font-size: 12px;
    color: #999;
}
</style>