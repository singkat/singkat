export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@unocss/nuxt',
    '@nuxt/image',
  ],
  css: [
    '@/public/style.css',
  ],
});
