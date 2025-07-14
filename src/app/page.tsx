'use client';

import React from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { FileText } from 'lucide-react';

export default function Home() {
  const router = useRouter();

  const actions = [
    {
      title: 'Neue Steuererklärung',
      description: 'Erstellen Sie eine neue Steuererklärung für das aktuelle Jahr',
      icon: <FileText className="h-8 w-8 text-blue-500" />,
      onClick: () => router.push('/steuererklaerung/neu')
    }
  ];

  return (
    <div className="container mx-auto px-4 py-8">
      <header className="mb-12 text-center">
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          Willkommen bei Baristeuer
        </h1>
        <p className="text-xl text-gray-600 dark:text-gray-300">
          Ihr zuverlässiger Begleiter für die Steuererklärung
        </p>
      </header>

      <div className="flex justify-center mb-12">
        <div className="max-w-md w-full">
          {actions.map((action, index) => (
            <div 
              key={index} 
              className="border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 rounded-xl p-8 hover:shadow-xl dark:hover:shadow-2xl transition-all duration-300 cursor-pointer transform hover:scale-105"
              onClick={action.onClick}
            >
              <div className="text-center mb-6">
                <div className="inline-flex p-4 bg-blue-50 dark:bg-blue-900/30 rounded-full mb-4">
                  {action.icon}
                </div>
                <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-2">{action.title}</h2>
                <p className="text-gray-600 dark:text-gray-300 leading-relaxed">
                  {action.description}
                </p>
              </div>
              <Button 
                className="w-full bg-blue-600 hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg transition-colors duration-200"
                size="lg"
              >
                Jetzt starten
              </Button>
            </div>
          ))}
        </div>
      </div>

      <footer className="mt-20 text-center text-gray-500 dark:text-gray-400 text-sm">
        <div className="border-t border-gray-200 dark:border-gray-700 pt-8">
          <p className="mb-4">© {new Date().getFullYear()} Baristeuer. Alle Rechte vorbehalten.</p>
          <div className="flex justify-center space-x-6">
            <a href="#" className="hover:text-blue-600 dark:hover:text-blue-400 transition-colors">Datenschutz</a>
            <a href="#" className="hover:text-blue-600 dark:hover:text-blue-400 transition-colors">Nutzungsbedingungen</a>
            <a href="#" className="hover:text-blue-600 dark:hover:text-blue-400 transition-colors">Impressum</a>
          </div>
        </div>
      </footer>
    </div>
  );
}
