<script>
  import { onMount } from 'svelte';
  import { Backend } from './wailsjs/go/service/DataService';

  let members = [];
  let name = '';
  let email = '';
  let joinDate = '';

  let editing = null;
  let editName = '';
  let editEmail = '';
  let editJoinDate = '';

  onMount(async () => {
    members = await Backend.ListMembers();
  });

  async function addMember() {
    if (!name || !email || !joinDate) return;
    await Backend.AddMember(name, email, joinDate);
    members = await Backend.ListMembers();
    name = '';
    email = '';
    joinDate = '';
  }

  function startEdit(m) {
    editing = m.id;
    editName = m.name;
    editEmail = m.email;
    editJoinDate = m.joinDate;
  }

  async function saveEdit(id) {
    await Backend.UpdateMember(id, editName, editEmail, editJoinDate);
    members = await Backend.ListMembers();
    editing = null;
  }

  function cancelEdit() {
    editing = null;
  }

  async function remove(id) {
    await Backend.DeleteMember(id);
    members = await Backend.ListMembers();
  }
</script>

<div>
  <h2 class="font-semibold mb-2">Mitglieder</h2>
  <div class="flex gap-2 mb-2">
    <input class="input input-bordered flex-1" placeholder="Name" bind:value={name} />
    <input class="input input-bordered flex-1" placeholder="Email" bind:value={email} />
    <input class="input input-bordered w-36" type="date" placeholder="Datum" bind:value={joinDate} />
    <button class="btn btn-primary" on:click={addMember}>Hinzufügen</button>
  </div>
  <table class="table w-full">
    <thead>
      <tr><th>Name</th><th>Email</th><th>Datum</th><th></th></tr>
    </thead>
    <tbody>
      {#each members as m}
        {#if editing === m.id}
          <tr>
            <td><input class="input input-bordered" bind:value={editName} /></td>
            <td><input class="input input-bordered" bind:value={editEmail} /></td>
            <td><input class="input input-bordered" type="date" bind:value={editJoinDate} /></td>
            <td class="flex gap-1">
              <button class="btn btn-primary btn-sm" on:click={() => saveEdit(m.id)}>Speichern</button>
              <button class="btn btn-secondary btn-sm" on:click={cancelEdit}>Abbrechen</button>
            </td>
          </tr>
        {:else}
          <tr>
            <td>{m.name}</td>
            <td>{m.email}</td>
            <td>{m.joinDate}</td>
            <td class="flex gap-1">
              <button class="btn btn-sm" on:click={() => startEdit(m)}>Bearbeiten</button>
              <button class="btn btn-error btn-sm" on:click={() => remove(m.id)}>Löschen</button>
            </td>
          </tr>
        {/if}
      {/each}
    </tbody>
  </table>
</div>
