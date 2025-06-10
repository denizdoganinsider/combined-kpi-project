<template>
  <div v-if="isAuthenticated" class="p-8">
    <h1 class="text-2xl font-bold text-gray-800 mb-4">API Response:</h1>
    <p class="text-lg text-green-600">{{ message }}</p>
  </div>
  <div v-else>
    <p>You're redirecting...</p>
  </div>
</template>

<script setup>
  definePageMeta({
    requiresAuth: true
  })

import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const isAuthenticated = process.client ? localStorage.getItem('auth_token') !== null : false;

const message = ref('Loading...')
const router = useRouter()

onMounted(async () => {
  if (!isAuthenticated) {
    router.push('/login')
  } else {
    try {
      const res = await axios.get('http://localhost:8080/api/v1/hello')
      message.value = res.data.message
    } catch (error) {
      message.value = 'Error occurred: ' + error.message
    }
  }
})
</script>

<style scoped>
</style>
