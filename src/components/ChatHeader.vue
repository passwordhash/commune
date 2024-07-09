<script setup>
const props = defineProps({
    user: {
        type: Object,
        default: {
            nickname: "Unkwonw"
        }
    }
})
</script>

<template>
    <div class="chat-header">
        <div class="wordwrap font-weight-bold"
        >{{user.nickname}}</div>
    </div>
</template>

<style scoped>
.chat-header {
    padding-left: 20px;
    font-size: 18px;
    font-weight: 800;
}

</style>