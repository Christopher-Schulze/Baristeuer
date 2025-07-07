<script>
  import { onMount } from 'svelte';
  import { Backend } from './wailsjs/go/service/DataService';
  import ProjectSelector from './ProjectSelector.svelte';
  let dark = false;
  let projectId = 0;
  let incomes = [];
  let incomeSource = '';
  let incomeAmount = '';
  let expenses = [];
  let expenseDesc = '';
  let expenseAmount = '';
  let showPDF = false;
  let pdfPath = '';
  let errorMsg = '';

  async function generatePdf() {
    errorMsg = '';
    pdfPath = '';
    try {
      pdfPath = await window.backend.Generator.GenerateReport(projectId);
      if (pdfPath) {
        window.open(`file://${pdfPath}`, '_blank');
      }
    } catch (err) {
      errorMsg = err?.message ?? String(err);
    }
  }

  onMount(async () => {
    const projects = await Backend.ListProjects();
    if (projects.length) {
      projectId = projects[0].id ?? projects[0].ID ?? projects[0].Id ?? 0;
      await loadData();
    }
  });

  async function loadData() {
    if (!projectId) return;
    incomes = await Backend.ListIncomes(projectId);
    expenses = await Backend.ListExpenses(projectId);
  }

  function handleSelect(event) {
    projectId = event.detail.id;
    loadData();
  }

  async function addIncome() {
    if (!incomeSource || !incomeAmount) return;
    await Backend.AddIncome(projectId, incomeSource, parseFloat(incomeAmount));
    incomes = await Backend.ListIncomes(projectId);
    incomeSource = '';
    incomeAmount = '';
  }

  async function addExpense() {
    if (!expenseDesc || !expenseAmount) return;
    await Backend.AddExpense(projectId, expenseDesc, parseFloat(expenseAmount));
    expenses = await Backend.ListExpenses(projectId);
    expenseDesc = '';
    expenseAmount = '';
  }
</script>

<main class="p-4" data-theme={dark ? 'dark' : 'light'}>
  <div class="flex justify-between mb-4">
    <h1 class="text-2xl font-bold">Bari$teuer</h1>
    <label class="flex items-center gap-2">
      <span>Dark</span>
      <input type="checkbox" bind:checked={dark} class="toggle" />
    </label>
  </div>
  <div class="tabs tabs-boxed mb-4">
    <a class="tab tab-active">Formulare</a>
    <a class="tab">Einstellungen</a>
    <a class="tab">PDF</a>
  </div>
  <ProjectSelector on:select={handleSelect} />
  <div class="grid md:grid-cols-2 gap-4">
    <div>
      <h2 class="font-semibold mb-2">Einnahmen</h2>
      <div class="flex gap-2 mb-2">
        <input class="input input-bordered flex-1" placeholder="Quelle" bind:value={incomeSource} />
        <input class="input input-bordered w-24" type="number" placeholder="Betrag" bind:value={incomeAmount} />
        <button class="btn btn-primary" on:click={addIncome}>Hinzufügen</button>
      </div>
      <table class="table w-full">
        <thead><tr><th>Quelle</th><th>Betrag</th></tr></thead>
        <tbody>
          {#each incomes as inc}
          <tr><td>{inc.source}</td><td>{inc.amount.toFixed(2)}</td></tr>
          {/each}
        </tbody>
      </table>
    </div>
    <div>
      <h2 class="font-semibold mb-2">Ausgaben</h2>
      <div class="flex gap-2 mb-2">
        <input class="input input-bordered flex-1" placeholder="Beschreibung" bind:value={expenseDesc} />
        <input class="input input-bordered w-24" type="number" placeholder="Betrag" bind:value={expenseAmount} />
        <button class="btn btn-primary" on:click={addExpense}>Hinzufügen</button>
      </div>
      <table class="table w-full">
        <thead><tr><th>Beschreibung</th><th>Betrag</th></tr></thead>
        <tbody>
          {#each expenses as ex}
          <tr><td>{ex.desc}</td><td>{ex.amount.toFixed(2)}</td></tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
  <div class="mt-6">
    <button class="btn" on:click={() => showPDF = !showPDF}>PDF Vorschau</button>
    {#if showPDF}
      <div class="mt-2 border p-2 flex flex-col gap-2" title="PDF Preview">
        <button class="btn btn-primary w-fit" on:click={generatePdf}>PDF erzeugen</button>
        {#if errorMsg}
          <p class="text-red-600">{errorMsg}</p>
        {/if}
        {#if pdfPath}
          <a class="link link-primary" href={`file://${pdfPath}`} download>Download</a>
        {/if}
      </div>
    {/if}
  </div>
</main>
