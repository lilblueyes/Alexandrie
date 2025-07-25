<template>
  <div v-if="showSearchModal" class="modal-mask">
    <div class="modal-container" ref="searchContainer">
      <span class="search-input">
        <Icon name="search" />
        <input type="text" v-model="filter" placeholder="Search document" ref="searchInput" />
        <button @click="close" style="background: none">
          <Icon name="close" :big="true" />
        </button>
      </span>
      <span class="title">Documents</span>
      <NuxtLink v-if="docs.length" v-for="doc of docs" class="item-search" :to="`/dashboard/docs/${doc.id}`" @click="close">
        <Icon name="draft" :big="true" fill="var(--font-color)" /> {{ doc.name }} <span v-html="categoryName(doc.category)" class="category"></span
      ></NuxtLink>
      <p v-else class="no-result">No result found.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
const documents = computed(() => useDocumentsStore().getAll.filter(c => c.name.toLowerCase().includes(filter.value.toLowerCase()) || c.tags?.toLowerCase().includes(filter.value.toLowerCase())));
const categoryName = (id: string = '') => {
  const category = useCategoriesStore().getById(id);
  if (category) return `<tag ${getAppColor(category.color)}>${category.name}</tag>`;
  return '';
};
const filter = ref<string>('');
const showSearchModal = ref<boolean>(false);

// Docs = max 5 docs
const docs = computed(() => documents.value.slice(0, 5));

const close = () => {
  showSearchModal.value = false;
  filter.value = '';
};
const searchContainer = ref<HTMLElement | null>(null);
const searchInput = ref<HTMLInputElement | null>(null);

const handleTab = (event: KeyboardEvent) => {
  event.preventDefault();
  if (!searchContainer.value) return;
  const focusableItems = Array.from(searchContainer.value.querySelectorAll('a'));
  const currentIndex = focusableItems.findIndex(item => item === document.activeElement);
  let nextIndex;
  if (event.shiftKey) {
    nextIndex = (currentIndex - 1 + focusableItems.length) % focusableItems.length;
  } else {
    nextIndex = (currentIndex + 1) % focusableItems.length;
  }
  focusableItems[nextIndex]?.focus();
};

onMounted(() => document.addEventListener('keydown', handleSearchShortCut));
onBeforeUnmount(() => document.removeEventListener('keydown', handleSearchShortCut));
const handleSearchShortCut = (e: KeyboardEvent) => {
  if (e.ctrlKey && e.key === 'q') {
    if (!showSearchModal.value) {
      showSearchModal.value = true;
      nextTick(() => searchInput.value?.focus());
    } else close();
  }
  if (showSearchModal.value) {
    if (e.key === 'Escape') close();
    if (e.key === 'Tab') handleTab(e);
  }
};
</script>

<style lang="scss" scoped>
.modal-container {
  width: 600px;
  margin: auto;
  padding: 20px;
  background-color: var(--bg-color);
  border-radius: 10px;
  box-shadow: 0 2px 8px var(--shadow);
  transition: all 0.3s ease;
}

.search-input {
  display: flex;
  align-items: center;

  input {
    background-color: var(--bg-color);
    border: none;
    outline: none;
    width: 100%;
  }

  svg {
    fill: var(--bg-contrast);
  }
}

.title {
  font-weight: 500;
  font-size: 0.9rem;
  color: #737373;
  padding: 10px 0;
}

.item-search {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin: 2.5px 0;
  padding: 10px 0;
  border-radius: 5px;
  color: var(--font-color);
  cursor: pointer;

  &:hover {
    background: var(--bg-contrast-2);
  }

  &:focus {
    outline: 2px solid var(--font-color);
  }

  .icon {
    display: flex;
    align-items: center;

    &:deep(svg) {
      width: 35px;
      margin-right: 5px;
    }
  }
}
.category {
  font-size: 1.1rem;
  color: var(--font-color);
  margin-left: auto;
  padding: 2px 5px;
  border-radius: 3px;
}

.no-result {
  text-align: center;
  font-weight: 500;
  font-size: 0.85rem;
}
</style>
