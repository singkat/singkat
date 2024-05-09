<script lang="ts" setup>
import { useToggle } from '@vueuse/core';

const [isLoading, setIsLoading] = useToggle();
const urlIn = ref('');
const urlOut = ref('');

const onSubmit = () => {
  if(!urlIn.value) return;

  setIsLoading();

  $fetch('http://localhost:3001/links', {
    method: 'POST',
    body: {
      url: urlIn.value,
    },
  }).then((res: any) => {
    urlOut.value = res?.data || '';
  }).finally(() => {
    setIsLoading();
  });
};

const onReset = () => {
  urlOut.value = '';
};
</script>

<template>
  <div class="flex flex-col mt-50 items-center h-[80vh]">
    <form @submit.prevent="onSubmit" class="w-1/2">
      <div class="flex">
        <input type="text" placeholder="Put your URL here!" v-model="urlIn" :disabled="isLoading" class="w-full py-2 px-3 rounded-l-full rounded-r-none" />
        <button type="submit" :disabled="isLoading" class="rounded-r-full rounded-l-none px-5">
          <span v-if="isLoading">Loading...</span>
          <span v-else>Singkat</span>
        </button>
      </div>
    </form>

    <div class="w-1/2 mt-3">
      <div v-if="!!urlOut" class="flex">
        <pre>{{ urlOut }}</pre>
        <button>copy</button>
        <button @click="onReset">reset</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
