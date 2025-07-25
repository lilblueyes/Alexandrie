<template>
  <div class="card-component">
    <h2 style="display: flex; align-items: center; justify-content: space-between">User Details <img style="border-radius: 50%; width: 40px; height: 40px" v-if="user" :src="useAvatar(user)" /></h2>
    <div v-if="user" style="width: 100%">
      <div class="user-details">
        <div class="user-detail">
          <label for="username">Username</label>
          <input id="username" v-model="user.username" type="text" />
        </div>
        <div class="user-detail">
          <label for="firstname">Firstname</label>
          <input id="firstname" v-model="user.firstname" type="text" />
        </div>
        <div class="user-detail">
          <label for="lastname">Lastname</label>
          <input id="lastname" v-model="user.lastname" type="text" />
        </div>
        <div class="user-detail">
          <label for="email">Email</label>
          <input id="email" v-model="user.email" type="email" />
        </div>
        <div class="user-detail">
          <label for="role">Role</label>
          <select id="role" v-model="user.role">
            <option :value="1">User</option>
            <option :value="2">Admin</option>
          </select>
        </div>
        <div style="display: flex; width: 100%">
          <div class="user-detail">
            <label>Created At</label>
            <div class="value">{{ new Date(user.created_timestamp).toLocaleDateString() }}</div>
          </div>
          <div class="user-detail">
            <label>Updated At</label>
            <div class="value">{{ new Date(user.updated_timestamp || 0).toLocaleDateString() }}</div>
          </div>
        </div>
      </div>

      <AppButton @click="saveChanges" type="success">Save Changes</AppButton>
      <AppButton @click="useRouter().push(`/dashboard/admin/users/${user.id}/documents`)" type="primary">View documents</AppButton>
    </div>
    <div v-else>No user found.</div>
  </div>
</template>

<script setup lang="ts">
import type { User } from '@/stores';

definePageMeta({ breadcrumb: 'User Details' });

const route = useRoute();
const user = ref<User | undefined>(undefined);
const store = useAdminStore();
watchEffect(async () => {
  user.value = await store.fetchById(route.params.id as string);
});

const saveChanges = async () => {
  if (!user.value) return;
  store
    .update(user.value)
    .then(() => useNotifications().add({ type: 'success', title: 'User updated successfully' }))
    .catch(e => useNotifications().add({ type: 'error', title: 'Error', message: e }));
};
</script>

<style scoped lang="scss">
.user-details {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 1rem;
  width: 100%;
}

.user-detail {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
  width: 100%;
}

label {
  font-weight: bold;
}
button {
  margin-top: 1rem;
}
</style>
