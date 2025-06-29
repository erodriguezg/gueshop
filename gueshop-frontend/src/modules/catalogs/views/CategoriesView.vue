<template>
    <div class="grid grid-cols-3 gap-4 ml-4 mt-4">
        <CatalogCard v-for="n in 10" :key="n" 
        :category="{ 
            name: `Categoria ${n}`, 
            description: `Very Long Description Number ${n}`, 
            imageUrl: 'https://www.fmdos.cl/wp-content/uploads/2016/03/joyas.jpg',
            slug: `category-n-${n}`}" />
    </div>
</template>

<script lang="ts" setup>
import CatalogCard from '../components/CategoryCard.vue';
</script>