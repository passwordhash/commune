<script setup>
const model = defineModel()

const props = defineProps({
  placeholder: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <div class="form-group">
    <input v-model="model" class="form-control" :placeholder="props.placeholder" required>
  </div>
</template>

<style scoped>

</style>