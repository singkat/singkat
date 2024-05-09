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
    <div class="flex flex-col">
      <div class="px-3 py-3 flex justify-between items-start hover:bg-slate-100">
        <div class="flex items-start space-x-3 w-11/12">
          <BAvatar src="https://www.jokesforfunny.com/wp-content/uploads/2021/06/0596bdb89b60fe771acd2f5972a9d3e3.jpg" size="sm" />
          <div class="flex flex-col w-3/4">
            <span>+1 Like</span>
            <span class="text-slate text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Eum id hic...</span>
          </div>
        </div>
        <div class="h-[10px] w-[10px] rounded-full bg-green"></div>
      </div>

      <div class="px-3 py-3 flex justify-between items-start hover:bg-slate-100">
        <div class="flex items-start space-x-3 w-11/12">
          <BAvatar src="https://www.jokesforfunny.com/wp-content/uploads/2021/06/0596bdb89b60fe771acd2f5972a9d3e3.jpg" size="sm" />
          <div class="flex flex-col">
            <span>New Comment from Alex</span>
            <span class="text-slate text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing</span>
          </div>
        </div>
        <div class="h-[10px] w-[10px] rounded-full bg-green"></div>
      </div>

    </div>
  </BSheetBottom>
</template>