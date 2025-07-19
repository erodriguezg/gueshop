<template>
    <div class="card bg-base-200 text-base-content w-72 shadow-md hover:scale-105 transition-transform duration-200">
        <figure>
            <img :src="props.product.imageUrl ?? notFoundUrl" alt="Imagen del producto"
                class="h-48 w-full object-cover" />
        </figure>
        <div class="card-body">
            <h2 class="card-title text-lg">{{ props.product.name }}</h2>
            <p class="text-sm text-gray-500 line-clamp-2">{{ props.product.shortDescription }}</p>
            <div class="mt-2 flex justify-between items-center">
                <span class="text-primary font-semibold text-lg">${{ props.product.price.toFixed(2) }}</span>
                <button class="btn btn-sm btn-primary">Ver</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const notFoundUrl = 'https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp'

import type { Product } from '../types/catalog'

const props = defineProps<{
    product: Product
}>()
</script>

<style scoped>
/* Optional: si usas Tailwind plugin line-clamp */
.line-clamp-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
}
</style>
