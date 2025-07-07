<script>
  import { onMount, createEventDispatcher } from 'svelte';
  import { Backend } from './wailsjs/go/service/DataService';

  const dispatch = createEventDispatcher();
  let projects = [];
  let newName = '';
  let selected = 0;

  onMount(loadProjects);

  async function loadProjects() {
    projects = await Backend.ListProjects();
    if (!selected && projects.length) {
      selected = projects[0].id ?? projects[0].ID ?? projects[0].Id ?? 0;
      dispatch('select', { id: selected });
    }
  }

  async function createProject() {
    if (!newName) return;
    const p = await Backend.CreateProject(newName);
    newName = '';
    await loadProjects();
    selected = p.id ?? p.ID ?? p.Id ?? 0;
    dispatch('select', { id: selected });
  }

  async function removeProject(id) {
    await Backend.DeleteProject(id);
    await loadProjects();
    if (selected === id) {
      selected = projects.length ? projects[0].id ?? projects[0].ID ?? projects[0].Id ?? 0 : 0;
      dispatch('select', { id: selected });
    }
  }

  function select(id) {
    selected = id;
    dispatch('select', { id });
  }
</script>

<div class="mb-4 border p-2 rounded">
  <h2 class="font-semibold mb-2">Projekte</h2>
  <div class="flex gap-2 mb-2">
    <input class="input input-bordered flex-1" placeholder="Projektname" bind:value={newName} />
    <button class="btn btn-primary" on:click={createProject}>Anlegen</button>
  </div>
  <ul class="space-y-1">
    {#each projects as proj}
      {#key proj.id}
      <li class="flex gap-2 items-center">
        <button class="btn btn-sm flex-1 {selected === (proj.id ?? proj.ID ?? proj.Id) ? 'btn-active' : ''}" on:click={() => select(proj.id ?? proj.ID ?? proj.Id)}>{proj.name}</button>
        <button class="btn btn-sm btn-error" on:click={() => removeProject(proj.id ?? proj.ID ?? proj.Id)}>✕</button>
      </li>
      {/key}
    {/each}
  </ul>
</div>
