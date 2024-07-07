<template>
  <div class="container">
    <h1 class="header">404</h1>
    <p class="message">Page not found</p>
  </div>
</template>
<script setup></script>
<style scoped>
.container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.header {
  font-size: 7rem;
  font-weight: bold;
  background: linear-gradient(to right, var(--green-light), var(--blue));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>
