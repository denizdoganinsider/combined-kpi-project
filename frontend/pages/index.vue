<template>
  <div v-if="isAuthenticated" class="p-8">
    <h1 class="text-2xl font-bold text-gray-800 mb-4">API Cevabı:</h1>
    <p class="text-lg text-green-600">{{ message }}</p>
  </div>
  <div v-else>
    <p>Yönlendiriliyorsunuz...</p>
  </div>
</template>

<script setup>
  definePageMeta({
    requiresAuth: true  // Bu sayfa giriş gerektiriyor
  })

import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

// Client-side kontrolü
const isAuthenticated = process.client ? localStorage.getItem('auth_token') !== null : false;

const message = ref('Yükleniyor...') // Başlangıç mesajı
const router = useRouter()

// Sayfa yüklendiğinde bu kod çalışacak
onMounted(async () => {
  if (!isAuthenticated) {
    // Eğer kullanıcı doğrulanmamışsa login sayfasına yönlendir
    router.push('/login')
  } else {
    // Kullanıcı doğrulandıysa API'den veri çek
    try {
      const res = await axios.get('http://localhost:8080/api/v1/hello')
      message.value = res.data.message // Gelen cevabı ekrana yazıyoruz
    } catch (error) {
      message.value = 'Hata oluştu: ' + error.message // Hata durumunda mesaj
    }
  }
})
</script>

<style scoped>
/* İstediğiniz stil düzenlemelerini buraya ekleyebilirsiniz */
</style>
