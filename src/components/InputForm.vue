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
.form-control {
    border-radius: 20px;
    border-color: #4CAF50;
    margin: 0 auto 10px;
    height: 50px;
    width: 100%;
}

.form-control:focus {
    border-color: #4CAF50;
    box-shadow: 0 0 0 0.2rem rgba(76, 175, 80, 0.25);
}
</style>