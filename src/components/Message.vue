<script setup>
import {onMounted} from "vue";

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
  if (props.isLast) {
    emit('mounted')
  }
})

</script>

<template>
  <div class="message d-block flex-column">
    <!--        <div><small class="text-muted">User1, 15:00</small></div>-->
    <div><small class="text-muted">{{ msg.date }}</small></div>
    <div class="message-text">{{ msg.text }}</div>
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
</style>