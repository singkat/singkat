<script lang="ts" setup>
import { onClickOutside } from '@vueuse/core';

withDefaults(defineProps<{
  show?: boolean,
}>(), {
  show: false,
});

const emits = defineEmits(['close']);

const sheetBottomRef = ref(null);
onClickOutside(sheetBottomRef, () => {
  emits('close');
});
</script>

<template>
  <Transition name="fade">
    <div v-if="show" class="fixed bottom-0 top-0 w-full opacity-30 bg-black"></div>
  </Transition>

  <Transition name="slide">
    <div v-if="show" ref="sheetBottomRef" class="fixed bg-white bottom-0 w-full pb-[5vh] rounded-t-3 transition">
      <div class="flex justify-center py-3">
        <div class="w-10 h-1 bg-slate-300 rounded-full"></div>
      </div>
  
      <slot />
    </div>
  </Transition>
</template>

<style>
</style>
