USE [cashier]
GO

/****** Object:  Table [dbo].[journals]    Script Date: 12/06/2022 08:29:34 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[journals](
	[journal_id] [int] IDENTITY(1,1) NOT NULL,
	[journal_date] [date] NOT NULL,
	[voucher_no] [varchar](50) NULL,
	[amount_beginning] [money] NULL,
	[amount_debit] [money] NULL,
	[amount_credit] [money] NULL,
	[amount_ending] [money] NULL,
	[description] [varchar](255) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
 CONSTRAINT [PK_journals] PRIMARY KEY CLUSTERED 
(
	[journal_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [dbo].[journals] ADD  CONSTRAINT [DF_journals_amount_beginning]  DEFAULT ((0)) FOR [amount_beginning]
GO

ALTER TABLE [dbo].[journals] ADD  CONSTRAINT [DF_journals_amount_debit]  DEFAULT ((0)) FOR [amount_debit]
GO

ALTER TABLE [dbo].[journals] ADD  CONSTRAINT [DF_journals_amount_credit]  DEFAULT ((0)) FOR [amount_credit]
GO

ALTER TABLE [dbo].[journals] ADD  CONSTRAINT [DF_journals_amount_ending]  DEFAULT ((0)) FOR [amount_ending]
GO


USE [cashier]
GO

/****** Object:  Table [dbo].[users]    Script Date: 12/06/2022 08:33:02 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[users](
	[users_id] [int] IDENTITY(1,1) NOT NULL,
	[email] [varchar](255) NOT NULL,
	[name] [varchar](255) NULL,
	[phone] [varchar](255) NULL,
	[address] [varchar](255) NULL,
	[ktp] [varchar](255) NULL,
 CONSTRAINT [PK_users] PRIMARY KEY CLUSTERED 
(
	[users_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO


