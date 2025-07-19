<template>
    <div class="grid grid-cols-4 gap-4 ml-4 mt-4">
        <ProductCard v-for="n in 20" :key="n" :product="{
            id: `12345-123456-${n}`,
            name: `Product Name ${n}`,
            imageUrl: 'https://picsum.photos/300',
            price: 999,
            shortDescription: `Short description of the product Number ${n}`
        }" />

    </div>
</template>

<script lang="ts" setup>
import ProductCard from '../components/ProductCard.vue';

</script>