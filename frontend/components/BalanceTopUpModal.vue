<template>
    <div
      v-if="open"
      class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
    >
      <div class="bg-white rounded-xl w-full max-w-sm p-6 text-black">
        <h2 class="text-lg font-semibold mb-4">Add Balance</h2>
  
        <input
          v-model.number="amount"
          type="number"
          min="1"
          placeholder="Amount"
          class="w-full border rounded px-3 py-2 mb-4 focus:outline-none focus:ring"
        />
  
        <p v-if="error" class="text-sm text-red-500 mb-3">
          {{ error }}
        </p>
  
        <div class="flex justify-end space-x-3">
          <button
            @click="$emit('close')"
            class="px-4 py-2 text-sm border rounded"
          >
            Cancel
          </button>
  
          <button
            @click="submit"
            :disabled="loading"
            class="px-4 py-2 text-sm bg-cyan-500 text-white rounded disabled:opacity-50"
          >
            {{ loading ? 'Adding...' : 'Add' }}
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useBalanceStore } from '@/stores/balance'
  
  defineProps({
    open: Boolean
  })
  
  const emit = defineEmits(['close'])
  
  const balanceStore = useBalanceStore()
  
  const amount = ref(null)
  const loading = ref(false)
  const error = ref(null)
  
  const submit = async () => {
    if (!amount.value || amount.value <= 0) {
      error.value = 'Please enter a valid amount'
      return
    }
  
    loading.value = true
    error.value = null
  
    try {
      await $fetch('/api/v1/balance/credit', {
        method: 'POST',
        body: {
          amount: amount.value
        },
        credentials: 'include'
      })
  
      await balanceStore.fetchCurrentBalance()
  
      emit('close')
      amount.value = null
    } catch (e) {
      error.value = 'Balance could not be added'
    } finally {
      loading.value = false
    }
  }
  </script>
  