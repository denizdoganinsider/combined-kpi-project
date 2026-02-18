export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()

  const rawPath = event.context.params!._
  const path = Array.isArray(rawPath) ? rawPath.join('/') : rawPath

  try {
    const token = getHeader(event, 'authorization')

    return await $fetch(`${config.public.apiBase}/api/v1/transactions/${path}`, {
      method: event.method,
      headers: {
        ...(token ? { Authorization: token } : {}),
      },
      body: ['POST', 'PUT', 'PATCH'].includes(event.method)
        ? await readBody(event)
        : undefined,
    })

  } catch (error: any) {

    throw createError({
      statusCode: error?.response?.status || 500,
      statusMessage: error?.response?._data?.error || 'Backend error'
    })
  }
})
