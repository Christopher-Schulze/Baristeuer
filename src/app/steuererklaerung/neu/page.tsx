import { SteuerFormular } from '@/components/steuer/SteuerFormular';

export default function NeueSteuererklaerung() {
  return (
    <div className="py-10">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900">Neue Steuererklärung</h1>
          <p className="mt-2 text-sm text-gray-600">
            Bitte füllen Sie das Formular aus, um Ihre Steuererklärung zu erstellen.
          </p>
        </div>
        
        <SteuerFormular vereinId="1" jahr={2024} />
      </div>
    </div>
  );
}
