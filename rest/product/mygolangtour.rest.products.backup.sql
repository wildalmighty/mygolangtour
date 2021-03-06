USE [mygolangtour.rest.products]
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Products]') AND type in (N'U'))
DROP TABLE [dbo].[Products]
GO
DROP USER [mygolangtour]
GO
USE [master]
GO
DROP DATABASE [mygolangtour.rest.products]
GO
CREATE DATABASE [mygolangtour.rest.products]

GO
ALTER DATABASE [mygolangtour.rest.products] SET COMPATIBILITY_LEVEL = 150
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [mygolangtour.rest.products].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
EXEC sys.sp_db_vardecimal_storage_format N'mygolangtour.rest.products', N'ON'
GO
ALTER DATABASE [mygolangtour.rest.products] SET QUERY_STORE = OFF
GO
USE [mygolangtour.rest.products]
GO
CREATE USER [mygolangtour] FOR LOGIN [mygolangtour] WITH DEFAULT_SCHEMA=[dbo]
GO
ALTER ROLE [db_datareader] ADD MEMBER [mygolangtour]
GO
ALTER ROLE [db_datawriter] ADD MEMBER [mygolangtour]
GO
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Products](
	[ProductID] [bigint] IDENTITY(1,1) NOT NULL,
	[Manufacturer] [nvarchar](50) NOT NULL,
	[ProductName] [nvarchar](50) NOT NULL,
 CONSTRAINT [PK_Products] PRIMARY KEY CLUSTERED 
(
	[ProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
SET IDENTITY_INSERT [dbo].[Products] ON 

INSERT [dbo].[Products] ([ProductID], [Manufacturer], [ProductName]) VALUES (1, N'Nike', N'Sky')
INSERT [dbo].[Products] ([ProductID], [Manufacturer], [ProductName]) VALUES (2, N'Addiddass', N'Earth')
INSERT [dbo].[Products] ([ProductID], [Manufacturer], [ProductName]) VALUES (3, N'Reebok', N'Razer')
INSERT [dbo].[Products] ([ProductID], [Manufacturer], [ProductName]) VALUES (4, N'Fila', N'Elephant')
SET IDENTITY_INSERT [dbo].[Products] OFF
GO
USE [master]
GO
ALTER DATABASE [mygolangtour.rest.products] SET  READ_WRITE 
GO
