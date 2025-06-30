<template>
    <div class="flex flex-col">

        <!-- Top Menu-->
        <top-menu />
        <main class="flex flex-row">
            <!-- ProductsFilter -->
            <products-filter />

            <!-- RouterView -->
            <router-view />
        </main>

    </div>
</template>

<script setup lang="ts">
import TopMenu from '@/modules/common/components/TopMenu.vue';
import ProductsFilter from '../components/ProductsFilter.vue';

</script>