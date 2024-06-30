<template>
  <div class="body">
    <div class="card">
      <h1>Login</h1>
      <input type="text" placeholder="username" />
      <input type="password" placeholder="password" />
      <ButtonPrimary class="login-btn" text="Login" post-icon="fa-solid fa-arrow-right" />
    </div>
  </div>
</template>
<script setup>
import ButtonPrimary from '@/components/ButtonPrimary.vue'
</script>
<style scoped>
.body {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.card {
  background-color: var(--grey-01);

  display: flex;
  flex-direction: column;
  align-items: center;

  padding: 1rem 2rem;
  border-radius: 0.8rem;
  box-shadow:
    rgba(0, 0, 0, 0.25) 0px 13px 27px -5px,
    rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
}

.card > * {
  margin: 1rem 0;
}

.login-btn {
  align-self: flex-end;
}

.card > input {
  width: 100%;

  border-color: #fdfcfc1a;
  background-color: #1d191673;
  color: white;

  border-radius: 0.75rem;
  border-width: 1px;
  border-style: solid;

  padding: 0.375rem 0.725rem;
  margin-top: 0.5ch;
  appearance: none;
  line-height: 1.5rem;

  font-size: 1rem;

  outline: 2px solid transparent;

  outline-offset: 2px;
}
.card input:hover {
  border-color: var(--grey-03);
}
.card input:focus {
  border-color: var(--primary);
}
</style>
