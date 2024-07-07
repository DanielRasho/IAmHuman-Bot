<template>
  <div class="body">
    <div class="card">
      <img src="../assets/dragon.png" alt="" />
      <h1>Hola!</h1>
      <p>
        Este link te dara acceso a la clase
        <span>{{ class_subject }}</span>
        seccion <span> {{ class_section }} </span> del server.
      </p>
      <ButtonPrimary
        class="login-btn"
        text="Authorize"
        :is-disabled="isButtonDisable"
        pre-icon="fa-brands fa-discord"
        @clicked="navigateToAuth"
      />
    </div>
  </div>
</template>
<script setup>
import ButtonPrimary from '@/components/ButtonPrimary.vue'
import { APISettings } from '@/utils/apiConfig'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const ouath_uri = ref('')
const class_subject = ref('...')
const class_section = ref('...')
const isButtonDisable = ref(true)

function navigateToAuth() {
  location.assign(ouath_uri.value)
}

onMounted(async () => {
  const url = APISettings.baseURL + 'assign-role/?invitation_id=' + route.params.id
  try {
    const response = await fetch(url, {
      method: 'GET'
    })
    const data = await response.json()
    if (!response.ok) throw Error(data.message)
    class_subject.value = data.invitation_data.class
    class_section.value = data.invitation_data.section
    ouath_uri.value = data.Oauth
    isButtonDisable.value = false
  } catch (err) {
    alert('Cannot fetch invitation data\n' + err)
  }
})
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
