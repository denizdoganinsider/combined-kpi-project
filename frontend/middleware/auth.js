export default defineNuxtRouteMiddleware((to, from) => {
  if (process.client) {
    const isAuthenticated = localStorage.getItem('auth_token') !== null;
  
    if (to.meta.requiresAuth && !isAuthenticated) {
      return navigateTo('/login');
    }
  }
})