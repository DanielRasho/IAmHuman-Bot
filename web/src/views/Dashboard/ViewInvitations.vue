<template>
  <section class="section section--view">
    <h1 class="header">Invitaciones Existentes</h1>
    <div>
      <SpinLoader v-if="isLoading" />
      <DataTable :headers="tableHeaders" :fetchData="fetchInvitations" />
    </div>
  </section>
</template>
<script setup>
import SpinLoader from '@/components/SpinLoader.vue'
import DataTable from '@/components/table/DataTable.vue'
import { APISettings } from '@/utils/apiConfig'
import { onMounted, ref } from 'vue'

const isLoading = ref(false)

const tableHeaders = ref([
  { label: 'ID', field: 'id', copy: false, width: '4ch', bold: true },
  { label: 'Clase', field: 'class', copy: false, width: '20ch', bold: true },
  { label: 'Seccion', field: 'section', copy: false, width: '8ch', bold: true },
  { label: 'Invitacion', field: 'invitation_url', copy: true, width: '20ch', bold: true },
  { label: 'Usos restantes', field: 'missing_uses', copy: false, width: '10ch', bold: true },
  { label: 'Creado', field: 'created_at', copy: false, width: '15ch', bold: true },
  { label: 'Server ID', field: 'server_id', copy: false, width: '18ch', bold: true },
  { label: 'Role ID', field: 'role_id', copy: false, width: '18ch', bold: true }
])

async function fetchInvitations(page, limit) {
  const url = APISettings.baseURL + `dashboard/?limit=${limit}&page=${page}`
  const token = sessionStorage.getItem('token')
  try {
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        Authorization: token
      }
    })
    const data = await response.json()
    if (response.ok == false) throw Error(data.message)
    return data
  } catch (err) {
    alert('Error\n' + err)
  }
}
</script>
<style scoped>
.header {
  color: var(--white);
  font-weight: bold;
  font-size: 2.2rem;
  margin-bottom: 1rem;
}
</style>
