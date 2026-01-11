export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const rawPath = event.context.params!._
    const path = Array.isArray(rawPath) ? rawPath.join('/') : rawPath

    return await $fetch(`${config.public.apiBase}/api/v1/transactions/${path}`, {
      method: event.method,
      body: ['POST', 'PUT', 'PATCH'].includes(event.method)
        ? await readBody(event)
        : undefined
    })
  })
  