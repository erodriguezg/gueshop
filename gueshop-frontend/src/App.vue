<template>
  <div :data-theme="theme">
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const theme = ref<string>("aqua")

</script>