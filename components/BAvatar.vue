<script lang="ts" setup>
const props = withDefaults(defineProps<{
  src: string,
  size?: 'sm' | 'md' | 'lg',
}>(), {
  size: 'md',
});
const sizeClass = computed(() => ({
  'sm': 'w-10 h-10',
  'md': 'w-15 h-15',
  'lg': 'w-20 h-20',
}[props.size]));
</script>

<template>
  <NuxtImg :class="sizeClass" class="rounded-full overflow-hidden object-cover" :src="src" />
</template>
