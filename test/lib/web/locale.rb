# frozen_string_literal: true

module Web
  # Expected translations for localized site responses in the acceptance suite.
  module Locale
    PAGES = {
      'de' => {
        'root' => {
          title: 'Lean Thoughts',
          description: 'Lean Thoughts teilt Ideen zu Lean Thinking und kontinuierlicher Verbesserung.',
          heading: 'Willkommen',
          quote: '„Perfektion ist unerreichbar, aber wenn wir nach Perfektion streben, können wir Spitzenleistungen erzielen.“',
          introduction: 'Lean Thinking ist letztlich eine Denkweise, eine Art, die Welt durch kontinuierliche Verbesserung zu betrachten.'
        },
        'books' => {
          title: 'Bücher | Lean Thoughts',
          description: 'Empfohlene Bücher über Lean Thinking, Softwarebereitstellung und den Arbeitsfluss in Teams.',
          heading: 'Bücher',
          quote: '„Wer heute liest, führt morgen.“',
          introduction: 'Lesen ist wichtig, weil es den Geist entwickelt und wertvolles Wissen und Lektionen fürs Leben vermittelt.'
        },
        'missing' => {
          title: 'Seite nicht gefunden | Lean Thoughts',
          description: 'Die angeforderte Seite von Lean Thoughts konnte nicht gefunden werden.',
          heading: 'Seite nicht gefunden',
          introduction: 'Die angeforderte Seite konnte nicht gefunden werden.'
        }
      },
      'es' => {
        'root' => {
          title: 'Lean Thoughts',
          description: 'Lean Thoughts comparte ideas sobre pensamiento lean y mejora continua.',
          heading: 'Bienvenido',
          quote: '«La perfección no se puede alcanzar, pero si la perseguimos podemos lograr la excelencia».',
          introduction: 'El pensamiento lean es, en esencia, una mentalidad, una forma de ver el mundo a través de la mejora continua.'
        },
        'books' => {
          title: 'Libros | Lean Thoughts',
          description: 'Libros recomendados sobre pensamiento lean, entrega de software y flujo de trabajo en equipo.',
          heading: 'Libros',
          quote: '«Hoy un lector, mañana un líder».',
          introduction: 'Leer es importante porque desarrolla la mente y aporta conocimientos y lecciones valiosas para la vida.'
        },
        'missing' => {
          title: 'Página no encontrada | Lean Thoughts',
          description: 'No se ha podido encontrar la página solicitada de Lean Thoughts.',
          heading: 'Página no encontrada',
          introduction: 'No se ha podido encontrar la página solicitada.'
        }
      }
    }.freeze

    LAYOUTS = {
      'de' => { home: 'Startseite', books: 'Bücher', code: 'Quellcode', contact: 'Kontakt aufnehmen', rights: 'Alle Rechte vorbehalten.' },
      'es' => { home: 'Inicio', books: 'Libros', code: 'Código', contact: 'Contactar', rights: 'Todos los derechos reservados.' }
    }.freeze

    BOOK_HEADERS = {
      'de' => %w[Titel Autor Themen Warum Link],
      'es' => ['Título', 'Autor', 'Temas', 'Por qué', 'Enlace']
    }.freeze

    BOOK_DETAILS = {
      'de' => [
        {
          topics: 'Kanban, evolutionäre Veränderung, Arbeitsfluss',
          note: 'Eine praktische Einführung in bessere Softwarebereitstellung durch visuelles Arbeitsmanagement und evolutionäre Veränderung.'
        },
        {
          topics: 'Softwareentwicklung, Continuous Delivery, Qualität',
          note: 'Verbindet technische Disziplin mit schnellem Feedback, Lernen und zuverlässiger Softwarebereitstellung.'
        },
        {
          topics: 'Teamgestaltung, kognitive Belastung, Arbeitsfluss',
          note: 'Erklärt, wie Teamstrukturen und Interaktionsmuster die Ergebnisse der Softwarebereitstellung prägen.'
        }
      ],
      'es' => [
        {
          topics: 'Kanban, cambio evolutivo, flujo',
          note: 'Una introducción práctica a la mejora de la entrega mediante la gestión visual del trabajo y el cambio evolutivo.'
        },
        {
          topics: 'ingeniería de software, entrega continua, calidad',
          note: 'Conecta la disciplina de ingeniería con la retroalimentación rápida, el aprendizaje y la entrega fiable de software.'
        },
        {
          topics: 'diseño de equipos, carga cognitiva, flujo',
          note: 'Explica cómo las estructuras de los equipos y sus formas de interacción influyen en la entrega de software.'
        }
      ]
    }.freeze
  end
end
