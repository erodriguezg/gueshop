<template>
    <div class="flex flex-col">
        <!-- Top Menu-->
        <top-menu />
        <main class="flex flex-row">
            <!-- RouterView -->
            <router-view />
        </main>

    </div>
</template>

<script setup lang="ts">
import TopMenu from '@/modules/common/components/TopMenu.vue';

</script>
