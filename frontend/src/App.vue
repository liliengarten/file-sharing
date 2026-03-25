<script setup>
import AppHeader from "@/components/AppHeader.vue";
import {useRoute} from "vue-router";
import {computed} from "vue";

const route = useRoute()

const displayHeader = computed(() => {
  return !["/login", "/register"].includes(route.path)
})
</script>

<template>
  <app-header v-if="displayHeader"></app-header>
  <router-view></router-view>
</template>

<style scoped></style>
