<script lang="ts" setup>
defineEmits(['close']);
withDefaults(defineProps<{
  show?: boolean,
}>(), {
  show: false,
});
</script>

<template>
  <BSheetBottom :show="show" @close="$emit('close')">
    <div class="px-5">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <BAvatar src="https://www.jokesforfunny.com/wp-content/uploads/2021/06/0596bdb89b60fe771acd2f5972a9d3e3.jpg" size="sm" />
          <div class="flex flex-col">
            <span class="font-bold">Hadi Hammurabi</span>
            <span>hadihammurabi@gmail.com</span>
          </div>
        </div>
        <BButton size="icon" variant="danger"><div class="i-uil-exit"></div></BButton>
      </div>
    </div>
  </BSheetBottom>
</template>
