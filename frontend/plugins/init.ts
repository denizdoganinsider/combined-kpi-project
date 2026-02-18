export default defineNuxtPlugin(() => {
    const apiStore = useApiStore()
    apiStore.init()
  })