<template>
    <div class="p-10 max-w-xl mx-auto">
      <h1 class="text-2xl font-bold mb-4">Wallet</h1>
  
      <div class="space-y-4">
        <div>
          <input v-model="creditAmount" class="border p-2 w-full" placeholder="Credit amount" />
          <button @click="handleCredit" class="btn">Credit</button>
        </div>
  
        <div>
          <input v-model="debitAmount" class="border p-2 w-full" placeholder="Debit amount" />
          <button @click="handleDebit" class="btn">Debit</button>
        </div>
  
        <div>
          <input v-model="toUser" class="border p-2 w-full" placeholder="To user id" />
          <input v-model="transferAmount" class="border p-2 w-full mt-2" placeholder="Amount" />
          <button @click="handleTransfer" class="btn">
            Transfer
          </button>
        </div>
      </div>
  
      <h2 class="mt-6 font-bold">Transactions</h2>
  
      <div v-if="store.loading" class="mt-2">Loading...</div>
      <div v-if="store.error" class="mt-2 text-red-500">{{ store.error }}</div>
  
      <table v-if="!store.loading" class="w-full border mt-2">
        <tr>
          <th>ID</th>
          <th>Type</th>
          <th>Amount</th>
          <th>Status</th>
        </tr>
        <tr v-for="t in store.transactions" :key="t.id">
          <td>{{ t.id }}</td>
          <td>{{ t.type }}</td>
          <td>{{ t.amount }}</td>
          <td>{{ t.status }}</td>
        </tr>
      </table>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useTransactionsStore } from '@/stores/transactions'
  
  const userId = 1
  const store = useTransactionsStore()
  
  const creditAmount = ref('')
  const debitAmount = ref('')
  const toUser = ref('')
  const transferAmount = ref('')
  
  onMounted(() => {
    store.fetchHistory(userId)
  })
  
  async function handleCredit() {
    await store.credit(userId, Number(creditAmount.value))
    creditAmount.value = ''
    await store.fetchHistory(userId)
  }
  
  async function handleDebit() {
    await store.debit(userId, Number(debitAmount.value))
    debitAmount.value = ''
    await store.fetchHistory(userId)
  }
  
  async function handleTransfer() {
    await store.transfer(
      userId,
      Number(toUser.value),
      Number(transferAmount.value)
    )
  
    toUser.value = ''
    transferAmount.value = ''
    await store.fetchHistory(userId)
  }
  </script>
  
  <style>
  .btn {
    background: black;
    color: white;
    padding: 8px;
    width: 100%;
    margin-top: 5px;
  }
  </style>
  