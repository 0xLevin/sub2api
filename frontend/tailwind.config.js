/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // 主色调 - 工业青绿
        primary: {
          50: '#effaf7',
          100: '#d6f3ec',
          200: '#aee7da',
          300: '#78d4c4',
          400: '#43bbaa',
          500: '#258f84',
          600: '#1f6f68',
          700: '#1c5954',
          800: '#1a4744',
          900: '#173b39',
          950: '#0a2221'
        },
        // 辅助色 - 石墨/铜色系
        accent: {
          50: '#f7f2ea',
          100: '#eee3d0',
          200: '#ddc69d',
          300: '#cda76d',
          400: '#bd8744',
          500: '#9b5f2f',
          600: '#7c4527',
          700: '#623522',
          800: '#4f2d20',
          900: '#42271d',
          950: '#24120c'
        },
        // 深色模式背景 - 暖石墨，避免原深蓝灰
        dark: {
          50: '#f4f0e8',
          100: '#e7ddca',
          200: '#d2c2a7',
          300: '#b8a484',
          400: '#918269',
          500: '#6f6655',
          600: '#4f4a3f',
          700: '#38362f',
          800: '#242520',
          900: '#191b18',
          950: '#101210'
        }
      },
      fontFamily: {
        sans: [
          'system-ui',
          '-apple-system',
          'BlinkMacSystemFont',
          'Segoe UI',
          'Roboto',
          'Helvetica Neue',
          'Arial',
          'PingFang SC',
          'Hiragino Sans GB',
          'Microsoft YaHei',
          'sans-serif'
        ],
        mono: ['ui-monospace', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', 'monospace']
      },
      boxShadow: {
        glass: '0 8px 32px rgba(0, 0, 0, 0.08)',
        'glass-sm': '0 4px 16px rgba(0, 0, 0, 0.06)',
        glow: '0 0 20px rgba(20, 184, 166, 0.25)',
        'glow-lg': '0 0 40px rgba(20, 184, 166, 0.35)',
        card: '0 1px 3px rgba(0, 0, 0, 0.04), 0 1px 2px rgba(0, 0, 0, 0.06)',
        'card-hover': '0 10px 40px rgba(0, 0, 0, 0.08)',
        'inner-glow': 'inset 0 1px 0 rgba(255, 255, 255, 0.1)'
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-primary': 'linear-gradient(135deg, #258f84 0%, #1f6f68 100%)',
        'gradient-dark': 'linear-gradient(135deg, #242520 0%, #101210 100%)',
        'gradient-glass':
          'linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 100%)',
        'mesh-gradient':
          'radial-gradient(at 40% 20%, rgba(31, 111, 104, 0.14) 0px, transparent 50%), radial-gradient(at 80% 0%, rgba(155, 95, 47, 0.10) 0px, transparent 50%), radial-gradient(at 0% 50%, rgba(67, 187, 170, 0.08) 0px, transparent 50%)'
      },
      animation: {
        'fade-in': 'fadeIn 0.3s ease-out',
        'slide-up': 'slideUp 0.3s ease-out',
        'slide-down': 'slideDown 0.3s ease-out',
        'slide-in-right': 'slideInRight 0.3s ease-out',
        'scale-in': 'scaleIn 0.2s ease-out',
        'pulse-slow': 'pulse 3s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        shimmer: 'shimmer 2s linear infinite',
        glow: 'glow 2s ease-in-out infinite alternate'
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' }
        },
        slideUp: {
          '0%': { opacity: '0', transform: 'translateY(10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' }
        },
        slideDown: {
          '0%': { opacity: '0', transform: 'translateY(-10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' }
        },
        slideInRight: {
          '0%': { opacity: '0', transform: 'translateX(20px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' }
        },
        scaleIn: {
          '0%': { opacity: '0', transform: 'scale(0.95)' },
          '100%': { opacity: '1', transform: 'scale(1)' }
        },
        shimmer: {
          '0%': { backgroundPosition: '-200% 0' },
          '100%': { backgroundPosition: '200% 0' }
        },
        glow: {
          '0%': { boxShadow: '0 0 20px rgba(31, 111, 104, 0.25)' },
          '100%': { boxShadow: '0 0 30px rgba(31, 111, 104, 0.4)' }
        }
      },
      backdropBlur: {
        xs: '2px'
      },
      borderRadius: {
        '4xl': '2rem'
      }
    }
  },
  plugins: []
}
