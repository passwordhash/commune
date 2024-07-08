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

<!--<template>-->
<!--    <div class="message d-block flex-column">-->
<!--        &lt;!&ndash;        <div><small class="text-muted">User1, 15:00</small></div>&ndash;&gt;-->
<!--        <div><small class="text-muted">{{ formatMessageTime }}</small></div>-->
<!--        <div class="message-text">{{ msg.text }}</div>-->
<!--    </div>-->
<!--</template>-->
<template>
    <div class="message-card">
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

.message {
    margin-bottom: 10px;
}

.own-message {
    text-align: right;
}

.message-text {
    display: inline-block;
    border-radius: 20px;
    padding: 10px;
    background-color: #FFFFFF;
    max-width: 70%;
    word-wrap: break-word;
}

.own-message .message-text {
    background-color: #4CAF50;
    color: #FFFFFF;
}

.message-card {
    background-color: #2C2F33;
    color: white;
    padding: 10px;
    border-radius: 15px;
    max-width: 300px;
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