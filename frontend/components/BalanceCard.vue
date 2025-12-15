<template>
    <div class="bg-gray-900 text-white rounded-xl p-4 w-full max-w-sm">
      <p class="text-sm text-gray-400">Current Balance</p>
  
      <div v-if="loading" class="mt-2 text-gray-500">
        Loading...
      </div>
  
      <div v-else class="mt-2 text-2xl font-semibold">
        {{ formattedBalance }}
      </div>
  
      <p v-if="error" class="text-sm text-red-400 mt-1">
        {{ error }}
      </p>
    </div>
  </template>
  
  <script setup>
  import { onMounted, computed } from 'vue'
  import { useBalanceStore } from '@/stores/balance'
  
  const balanceStore = useBalanceStore()
  
  onMounted(() => {
    balanceStore.fetchCurrentBalance()
  })
  
  const loading = computed(() => balanceStore.loading)
  const error = computed(() => balanceStore.error)
  
  const formattedBalance = computed(() => {
    if (balanceStore.amount === null) return '--'
    return `${balanceStore.amount.toLocaleString()} ₺`
  })
  </script>
  