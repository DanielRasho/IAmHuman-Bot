<template>
  <div class="body">
    <div class="card">
      <img src="../assets/dragon.png" alt="" />
      <h1>Hola!</h1>
      <p>
        Este link te dara acceso a la clase
        <span>POO</span>
        seccion <span> 10 </span> del server.
      </p>
      <p>Usos restantes: <span> 15</span></p>
      <div class="auth-container">
        <ButtonPrimary class="login-btn" text="Authorize" pre-icon="fa-brands fa-discord" />
        <ToolTip
          message="We need your permission to access to your account ID"
          width="30ch"
          style="margin-left: 2ch"
        />
      </div>
    </div>
  </div>
</template>
<script setup>
import ButtonPrimary from '@/components/ButtonPrimary.vue'
import ToolTip from '@/components/ToolTip.vue'
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
  width: 40%;
  max-width: 600px;
  min-width: 400px;

  padding: 1rem 2rem;
  border-radius: 0.8rem;
  box-shadow:
    rgba(0, 0, 0, 0.25) 0px 13px 27px -5px,
    rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
}

.card > * {
  margin: 1rem 0;
}

.card > h1 {
  font-weight: bolder;
  font-size: 3.5rem;
  letter-spacing: 0.1ch;
  background: linear-gradient(to right, var(--green-light), var(--blue));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.card > img {
  border-radius: 50%;
  width: 15ch;
}

.card > p {
  text-align: center;
}

.card > p > span {
  font-weight: bold;
  color: var(--white);
}

@media only screen and (max-width: 700px) {
  .card {
    background-color: transparent;
    box-shadow: none;
    width: 100%;
  }
}
</style>
