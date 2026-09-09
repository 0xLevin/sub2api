export default {
  batchImageGuide: {
    title: 'Batch Image Generation',
    description: 'Submit multiple prompts in one job and download the generated images when complete'
  },
  // Home Page
  home: {
    viewOnGithub: 'View on GitHub',
    viewDocs: 'View Documentation',
    docs: 'Docs',
    switchToLight: 'Switch to Light Mode',
    switchToDark: 'Switch to Dark Mode',
    dashboard: 'Dashboard',
    login: 'Login',
    getStarted: 'Get Started',
    goToDashboard: 'Go to Dashboard',
    // User-focused value proposition
    heroSubtitle: 'One Key, All AI Models',
    heroDescription: 'No need to manage multiple subscriptions. Access Claude, GPT, Gemini and more with a single API key',
    tags: {
      subscriptionToApi: 'Subscription to API',
      stickySession: 'Session Persistence',
      realtimeBilling: 'Pay As You Go'
    },
    // Pain points section
    painPoints: {
      title: 'Sound Familiar?',
      items: {
        expensive: {
          title: 'High Subscription Costs',
          desc: 'Paying for multiple AI subscriptions that add up every month'
        },
        complex: {
          title: 'Account Chaos',
          desc: 'Managing scattered accounts and API keys across different platforms'
        },
        unstable: {
          title: 'Service Interruptions',
          desc: 'Single accounts hitting rate limits and disrupting your workflow'
        },
        noControl: {
          title: 'No Usage Control',
          desc: "Can't track where your money goes or limit team member usage"
        }
      }
    },
    // Solutions section
    solutions: {
      title: 'We Solve These Problems',
      subtitle: 'Three simple steps to stress-free AI access'
    },
    features: {
      unifiedGateway: 'One-Click Access',
      unifiedGatewayDesc: 'Get a single API key to call all connected AI models. No separate applications needed.',
      multiAccount: 'Always Reliable',
      multiAccountDesc: 'Smart routing across multiple upstream accounts with automatic failover. Say goodbye to errors.',
      balanceQuota: 'Pay What You Use',
      balanceQuotaDesc: 'Usage-based billing with quota limits. Full visibility into team consumption.'
    },
    // Comparison section
    comparison: {
      title: 'Why Choose Us?',
      headers: {
        feature: 'Comparison',
        official: 'Official Subscriptions',
        us: 'Our Platform'
      },
      items: {
        pricing: {
          feature: 'Pricing',
          official: 'Fixed monthly fee, pay even if unused',
          us: 'Pay only for what you use'
        },
        models: {
          feature: 'Model Selection',
          official: 'Single provider only',
          us: 'Switch between models freely'
        },
        management: {
          feature: 'Account Management',
          official: 'Manage each service separately',
          us: 'Unified key, one dashboard'
        },
        stability: {
          feature: 'Stability',
          official: 'Single account rate limits',
          us: 'Multi-account pool, auto-failover'
        },
        control: {
          feature: 'Usage Control',
          official: 'Not available',
          us: 'Quotas & detailed analytics'
        }
      }
    },
    providers: {
      title: 'Supported AI Models',
      description: 'One API, Multiple Choices',
      supported: 'Supported',
      soon: 'Soon',
      claude: 'Claude',
      gemini: 'Gemini',
      antigravity: 'API',
      more: 'More'
    },
    // CTA section
    cta: {
      title: 'Ready to Get Started?',
      description: 'Sign up now and get free trial credits to experience seamless AI access',
      button: 'Sign Up Free'
    },
    v2: {
      kicker: 'OPENAI API ACCESS',
      buying: {
        kicker: 'Buy and Connect',
        title: 'Buy OpenAI tokens like cloud resources',
        description: 'Top up, generate an API key, and review balance, usage, and request history from one place.',
        items: {
          topUp: {
            label: 'Step 01',
            title: 'Top Up Balance',
            description: 'Choose a recharge amount and start calling after balance is credited, without decoding complex plans first.'
          },
          key: {
            label: 'Step 02',
            title: 'Get an API Key',
            description: 'Create keys for personal projects, teammates, or business lines so access is easy to rotate or disable.'
          },
          usage: {
            label: 'Step 03',
            title: 'Pay by Usage',
            description: 'Each request deducts balance and keeps a record, so customers can review GPT calls, tokens, and costs.'
          }
        }
      },
      capabilities: {
        kicker: 'What You Get',
        title: 'OpenAI tokens for GPT access',
        description: 'Keep the path from purchase to request simple: top up, get a key, connect your app, and track balance changes.',
        items: {
          api: {
            title: 'Compatible API',
            description: 'Use familiar OpenAI-compatible request formats so scripts, plugins, and product backends are easier to connect.'
          },
          routing: {
            title: 'GPT Workloads',
            description: 'Use OpenAI tokens for text generation, coding assistance, and automation calls.'
          },
          accountPool: {
            title: 'Top Up and Use',
            description: 'Buy tokens when you need them instead of maintaining multiple fixed subscriptions.'
          },
          billing: {
            title: 'Realtime Deduction',
            description: 'Every request records token usage, cost, and balance so spending stays understandable.'
          },
          guardrails: {
            title: 'Quota Controls',
            description: 'Set limits for API keys to avoid accidental overspending from scripts or teammates.'
          },
          observability: {
            title: 'Usage Lookup',
            description: 'Review calls by time, GPT usage, and API key for reconciliation and troubleshooting.'
          }
        }
      },
      workflow: {
        kicker: 'How It Works',
        title: 'Start spending tokens in four steps',
        description: 'From signup to your first API call, the path stays simple so you can focus on your app and prompts.',
        steps: {
          connect: {
            title: 'Create Account',
            description: 'Enter the dashboard and find balance, keys, orders, and usage in one place.'
          },
          issue: {
            title: 'Top Up Tokens',
            description: 'Buy OpenAI token credits as needed instead of carrying a fixed subscription for occasional calls.'
          },
          route: {
            title: 'Generate API Key',
            description: 'Create separate keys for projects, scripts, or teammates so access is easy to manage.'
          },
          measure: {
            title: 'Start Calling',
            description: 'Place the key in your app, spend tokens per request, and review details whenever needed.'
          }
        }
      },
      controlPlane: {
        kicker: 'Account View',
        metrics: {
          access: 'Access',
          models: 'Provider',
          usage: 'Usage',
          balance: 'Balance'
        },
        signals: {
          failover: 'Use one API key for OpenAI-compatible calls',
          usage: 'Record requests, tokens, and costs in realtime',
          audit: 'Split keys by project or teammate for clearer control'
        }
      },
      modelAccess: {
        kicker: 'OpenAI Access',
        title: 'One balance for common GPT workloads',
        description: 'Manage OpenAI token purchase, API keys, and usage records in one place for routine GPT calls.',
        items: {
          chat: {
            title: 'Chat and Content',
            description: 'Use it for support replies, drafts, document cleanup, knowledge base Q&A, and other text-heavy work.'
          },
          code: {
            title: 'Code and Dev Tools',
            description: 'Use it for script generation, code explanation, test help, IDE plugins, and internal developer assistants.'
          },
          automation: {
            title: 'Automation Workflows',
            description: 'Use it in scheduled jobs, bots, data pipelines, and product backends that need GPT calls.'
          }
        }
      },
      useCases: {
        kicker: 'Use Cases',
        title: 'For users who need reliable OpenAI token access',
        items: {
          teams: {
            title: 'Team Workflows',
            description: 'Give each member or project a key and track their token consumption.'
          },
          products: {
            title: 'Product Integration',
            description: 'Connect GPT capabilities to SaaS apps, bots, plugins, and automation workflows.'
          },
          resellers: {
            title: 'Individual Developers',
            description: 'Use one balance for experiments, development, and daily API calls without multiple subscriptions.'
          }
        }
      },
      faq: {
        kicker: 'FAQ',
        title: 'What to confirm before top-up',
        items: {
          billing: {
            question: 'How are tokens billed?',
            answer: 'Balance is deducted by actual request usage, and usage records are kept for lookup and reconciliation.'
          },
          compatibility: {
            question: 'Can existing apps connect?',
            answer: 'Use compatible APIs and API keys for scripts, plugins, backend services, and automation tools.'
          },
          balance: {
            question: 'Can I see balance details?',
            answer: 'Yes. Review request records, token usage, cost, and balance changes so spending is not a black box.'
          },
          keys: {
            question: 'Can projects use separate keys?',
            answer: 'Yes. Create separate keys for projects, members, or environments so usage is easier to limit and trace.'
          }
        }
      },
      finalCta: {
        kicker: 'Start Using'
      }
    },
    footer: {
      allRightsReserved: 'All rights reserved.'
    }
  },

  // Key Usage Query Page
  keyUsage: {
    title: 'API Key Usage',
    subtitle: 'Enter your API Key to view real-time spending and usage status',
    placeholder: 'sk-ant-mirror-xxxxxxxxxxxx',
    query: 'Query',
    querying: 'Querying...',
    privacyNote: 'Your Key is processed locally in the browser and will not be stored',
    dateRange: 'Date Range:',
    dateRangeToday: 'Today',
    dateRange7d: '7 Days',
    dateRange30d: '30 Days',
    dateRange90d: '90 Days',
    dateRangeCustom: 'Custom',
    apply: 'Apply',
    used: 'Used',
    detailInfo: 'Detail Information',
    tokenStats: 'Token Statistics',
    dailyDetail: 'Daily Detail',
    modelStats: 'Model Usage Statistics',
    // Table headers
    date: 'Date',
    model: 'Model',
    requests: 'Requests',
    inputTokens: 'Input Tokens',
    outputTokens: 'Output Tokens',
    cacheCreationTokens: 'Cache Creation',
    cacheReadTokens: 'Cache Read',
    cacheWriteTokens: 'Cache Write',
    totalTokens: 'Total Tokens',
    cost: 'Cost',
    // Status
    quotaMode: 'Key Quota Mode',
    walletBalance: 'Wallet Balance',
    // Ring card titles
    totalQuota: 'Total Quota',
    limit5h: '5-Hour Limit',
    limitDaily: 'Daily Limit',
    limit7d: '7-Day Limit',
    limitWeekly: 'Weekly Limit',
    limitMonthly: 'Monthly Limit',
    // Detail rows
    remainingQuota: 'Remaining Quota',
    expiresAt: 'Expires At',
    todayExpires: '(expires today)',
    daysLeft: '({days} days)',
    usedQuota: 'Used Quota',
    resetNow: 'Resetting soon',
    subscriptionType: 'Subscription Type',
    subscriptionExpires: 'Subscription Expires',
    // Usage stat cells
    todayRequests: 'Today Requests',
    todayInputTokens: 'Today Input',
    todayOutputTokens: 'Today Output',
    todayTokens: 'Today Tokens',
    todayCacheCreation: 'Today Cache Creation',
    todayCacheRead: 'Today Cache Read',
    todayCost: 'Today Cost',
    rpmTpm: 'RPM / TPM',
    totalRequests: 'Total Requests',
    totalInputTokens: 'Total Input',
    totalOutputTokens: 'Total Output',
    totalTokensLabel: 'Total Tokens',
    totalCacheCreation: 'Total Cache Creation',
    totalCacheRead: 'Total Cache Read',
    totalCost: 'Total Cost',
    avgDuration: 'Avg Duration',
    // Messages
    enterApiKey: 'Please enter an API Key',
    querySuccess: 'Query successful',
    queryFailed: 'Query failed',
    queryFailedRetry: 'Query failed, please try again later',
    noDailyUsage: 'No daily usage data',
  },

  // Setup Wizard
  setup: {
    title: 'Sub2API Setup',
    description: 'Configure your Sub2API instance',
    database: {
      title: 'Database Configuration',
      description: 'Connect to your PostgreSQL database',
      host: 'Host',
      port: 'Port',
      username: 'Username',
      password: 'Password',
      databaseName: 'Database Name',
      sslMode: 'SSL Mode',
      passwordPlaceholder: 'Password',
      ssl: {
        disable: 'Disable',
        require: 'Require',
        verifyCa: 'Verify CA',
        verifyFull: 'Verify Full'
      }
    },
    redis: {
      title: 'Redis Configuration',
      description: 'Connect to your Redis server',
      host: 'Host',
      port: 'Port',
      username: 'Username (optional)',
      password: 'Password (optional)',
      database: 'Database',
      usernamePlaceholder: 'Leave empty for default user',
      passwordPlaceholder: 'Password',
      enableTls: 'Enable TLS',
      enableTlsHint: 'Use TLS when connecting to Redis (public CA certs)'
    },
    admin: {
      title: 'Admin Account',
      description: 'Create your administrator account',
      email: 'Email',
      password: 'Password',
      confirmPassword: 'Confirm Password',
      passwordPlaceholder: 'Min 8 characters',
      confirmPasswordPlaceholder: 'Confirm password',
      passwordMismatch: 'Passwords do not match'
    },
    ready: {
      title: 'Ready to Install',
      description: 'Review your configuration and complete setup',
      database: 'Database',
      redis: 'Redis',
      adminEmail: 'Admin Email'
    },
    status: {
      testing: 'Testing...',
      success: 'Connection Successful',
      testConnection: 'Test Connection',
      installing: 'Installing...',
      completeInstallation: 'Complete Installation',
      completed: 'Installation completed!',
      redirecting: 'Redirecting to login page...',
      restarting: 'Service is restarting, please wait...',
      timeout: 'Service restart is taking longer than expected. Please refresh the page manually.'
    }
  },

  // Common
}
