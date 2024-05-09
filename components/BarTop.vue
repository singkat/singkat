<script lang="ts" setup>
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core';

defineEmits(['profileClick', 'notifClick']);

const ww = useBreakpoints(breakpointsTailwind);
</script>

<template>
  <header class="shadow-sm flex justify-between items-center px-5 py-2.5">
    <div>
      <NuxtLink to="/" class="no-underline text-3xl">
        W
      </NuxtLink>
    </div>
    <div class="flex space-x-3">
      <input type="text" class="form-input" v-if="ww.lg.value" placeholder="Search...">
      <BButton size="icon" v-else><div class="i-uil-search"></div></BButton>

      <BButton size="icon" @click="$emit('notifClick')"><div class="i-uil-bell"></div></BButton>

      <BAvatar src="https://www.jokesforfunny.com/wp-content/uploads/2021/06/0596bdb89b60fe771acd2f5972a9d3e3.jpg" size="sm" @click="$emit('profileClick')" />
    </div>
  </header>
</template>
