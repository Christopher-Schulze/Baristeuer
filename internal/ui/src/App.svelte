<script>
  import { onMount } from 'svelte';
  import { Backend } from './wailsjs/go/service/DataService';
  import { locale, _ } from './i18n';
  let dark = false;
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
      // Currently uses project ID 1; adapt as needed once projects are selectable
      pdfPath = await window.backend.Generator.GenerateReport(1);
      if (pdfPath) {
        window.open(`file://${pdfPath}`, '_blank');
      }
    } catch (err) {
      errorMsg = err?.message ?? String(err);
    }
  }

  let projectId = 0;

  onMount(async () => {
    const proj = await Backend.CreateProject('Default');
    projectId = proj.id ?? proj.ID ?? proj.Id ?? 0;
    incomes = await Backend.ListIncomes(projectId);
    expenses = await Backend.ListExpenses(projectId);
  });

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
    <h1 class="text-2xl font-bold">{$_('appTitle')}</h1>
    <div class="flex items-center gap-4">
      <label class="flex items-center gap-2">
        <span>{$_('dark')}</span>
        <input type="checkbox" bind:checked={dark} class="toggle" />
      </label>
      <select class="select select-bordered select-sm" bind:value={$locale}>
        <option value="de">DE</option>
        <option value="en">EN</option>
      </select>
    </div>
  </div>
  <div class="tabs tabs-boxed mb-4">
    <a class="tab tab-active">{$_('tab.forms')}</a>
    <a class="tab">{$_('tab.settings')}</a>
    <a class="tab">{$_('tab.pdf')}</a>
  </div>
  <div class="grid md:grid-cols-2 gap-4">
    <div>
      <h2 class="font-semibold mb-2">{$_('income.title')}</h2>
      <div class="flex gap-2 mb-2">
        <input class="input input-bordered flex-1" placeholder={$_('income.source')} bind:value={incomeSource} />
        <input class="input input-bordered w-24" type="number" placeholder={$_('income.amount')} bind:value={incomeAmount} />
        <button class="btn btn-primary" on:click={addIncome}>{$_('income.add')}</button>
      </div>
      <table class="table w-full">
        <thead><tr><th>{$_('income.source')}</th><th>{$_('income.amount')}</th></tr></thead>
        <tbody>
          {#each incomes as inc}
          <tr><td>{inc.source}</td><td>{inc.amount.toFixed(2)}</td></tr>
          {/each}
        </tbody>
      </table>
    </div>
    <div>
      <h2 class="font-semibold mb-2">{$_('expenses.title')}</h2>
      <div class="flex gap-2 mb-2">
        <input class="input input-bordered flex-1" placeholder={$_('expenses.desc')} bind:value={expenseDesc} />
        <input class="input input-bordered w-24" type="number" placeholder={$_('expenses.amount')} bind:value={expenseAmount} />
        <button class="btn btn-primary" on:click={addExpense}>{$_('expenses.add')}</button>
      </div>
      <table class="table w-full">
        <thead><tr><th>{$_('expenses.desc')}</th><th>{$_('expenses.amount')}</th></tr></thead>
        <tbody>
          {#each expenses as ex}
          <tr><td>{ex.desc}</td><td>{ex.amount.toFixed(2)}</td></tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
  <div class="mt-6">
    <button class="btn" on:click={() => showPDF = !showPDF}>{$_('pdf.preview')}</button>
    {#if showPDF}
      <div class="mt-2 border p-2 flex flex-col gap-2" title={$_('pdf.preview')}>
        <button class="btn btn-primary w-fit" on:click={generatePdf}>{$_('pdf.generate')}</button>
        {#if errorMsg}
          <p class="text-red-600">{errorMsg}</p>
        {/if}
        {#if pdfPath}
          <a class="link link-primary" href={`file://${pdfPath}`} download>{$_('download')}</a>
        {/if}
      </div>
    {/if}
  </div>
</main>
