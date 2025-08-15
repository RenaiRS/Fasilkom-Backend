-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Aug 15, 2025 at 03:38 AM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `fasilkom_be`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `token` longtext,
  `token_exp` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `token`, `token_exp`) VALUES
(1, '2025-08-13 19:09:46.152', '2025-08-15 02:09:04.577', NULL, 'admin web', '$2a$10$sNG.GstmomZZdN4Msu9mEuZAUAsmH.EWjzFi4ec2WlN4Xa64vVM.m', 'dcee385e8dc27150c73d6c3e4e9420978609f07ec7a7a6bf23875e0c75b8a4df', '2025-08-15 03:09:04.577');

-- --------------------------------------------------------

--
-- Table structure for table `berita`
--

CREATE TABLE `berita` (
  `id` bigint UNSIGNED NOT NULL,
  `judul` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `tanggal` date DEFAULT NULL,
  `cover` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `is_priority` tinyint(1) DEFAULT NULL,
  `posted_by` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `berita`
--

INSERT INTO `berita` (`id`, `judul`, `tanggal`, `cover`, `content`, `is_priority`, `posted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
(8, 'Kunjungan Fakultas Ilmu Komputer Universitas Sriwijaya ke IPB University', '2025-07-05', 'http://localhost:8080/assets/covers/1755198547868182300.png', 'Sebagai bagian dari komitmen untuk mencetak lulusan yang kompeten, adaptif, dan siap bersaing di tingkat global, Fakultas Ilmu Komputer Universitas Sriwijaya menyelenggarakan Program Pelatihan dan Sertifikasi Internasional Batch 1.', 1, 'admin web', '2025-08-15 02:09:07.871', '2025-08-15 07:12:49.368', NULL),
(9, 'Fakultas Ilmu Komputer Universitas Sriwijaya kembali melepas alumni baru / Yudisium Ke-75.', '2025-04-17', 'http://localhost:8080/assets/covers/1755198569955333700.jpg', '', 1, 'admin web', '2025-08-15 02:09:29.956', '2025-08-15 07:12:49.375', NULL),
(10, 'Selamat Hari Raya Idul Fitri 1446 H', '2025-03-31', 'http://localhost:8080/assets/covers/1755198583748987900.jpg', '', 1, 'admin web', '2025-08-15 02:09:43.750', '2025-08-15 07:12:49.378', NULL),
(11, 'Dies Natalis ke-19 Fasilkom Unsri', '2025-08-15', 'http://localhost:8080/assets/covers/1755198611616591700.jpg', 'Dalam rangka Dies Natalis ke-19, Fasilkom Unsri mengadakan beberapa rangkaian kegiatan yang diselenggarakan di Highland Resort, Bogor, Jawa Barat dari tanggal 28 Juli 2025 hingga 1 Agustus 2025.', 0, 'admin web', '2025-08-15 02:10:11.619', '2025-08-15 07:12:49.380', NULL),
(12, 'Fakultas Ilmu Komputer Universitas Sriwijaya menyelenggarakan kegiatan pelatihan bertajuk “Detox Emotion”', '2025-08-15', 'http://localhost:8080/assets/covers/1755198633671029200.jpg', 'Fakultas Ilmu Komputer Universitas Sriwijaya menyelenggarakan kegiatan pelatihan bertajuk “Detox Emotion: Mewujudkan Kesehatan dan Kesejahteraan Pegawai ASN Universitas Sriwijaya sebagai Wujud Dukungan Pembangunan Berkelanjutan” pada Sabtu, 12 Juli 2025. Kegiatan ini berlangsung di Hall Dr. Jaidan Jauhari, S.Pd., M.T., Lantai 7 Gedung Diklat Fasilkom Unsri Kampus Palembang.', 0, 'admin web', '2025-08-15 02:10:33.675', '2025-08-15 07:12:49.382', NULL),
(13, 'PENGUMUMAN ! PENGISIAN KRS SEMESTER GANJIL 2025/2026', '2025-08-15', 'http://localhost:8080/assets/covers/1755198663040298800.jpeg', 'Pengisian Kartu Studi Mahasiswa lama: 4 - 8 Agustus 2025\nPengisian Kartu Studi Mahasiswa baru: 15- 18 Agustus 2025\nPengisian Kartu Perubahan Studi Mahasiswa: 18 -22 Agustus 2025', 0, 'admin web', '2025-08-15 02:11:03.042', '2025-08-15 07:12:49.383', NULL),
(14, 'Peningkatan Kompetensi Mahasiswa Melalui Program Pelatihan dan Sertifikasi Internasional di FASILKOM UNSRI', '2025-08-15', 'http://localhost:8080/assets/covers/1755198684426246500.png', 'Sebagai bagian dari komitmen untuk mencetak lulusan yang kompeten, adaptif, dan siap bersaing di tingkat global, Fakultas Ilmu Komputer Universitas Sriwijaya menyelenggarakan Program Pelatihan dan Sertifikasi Internasional Batch 1.\n', 0, 'admin web', '2025-08-15 02:11:24.427', '2025-08-15 07:12:49.384', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `files`
--

CREATE TABLE `files` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` longtext,
  `file_path` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `kemahasiswaan_kerjasamas`
--

CREATE TABLE `kemahasiswaan_kerjasamas` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `judul` longtext NOT NULL,
  `tanggal` datetime(3) DEFAULT NULL,
  `cover` longtext,
  `content` longtext,
  `posted_by` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `kemahasiswaan_kerjasamas`
--

INSERT INTO `kemahasiswaan_kerjasamas` (`id`, `created_at`, `updated_at`, `deleted_at`, `judul`, `tanggal`, `cover`, `content`, `posted_by`) VALUES
(6, '2025-08-15 02:14:07.000', '2025-08-15 07:12:49.392', NULL, 'Registrasi Acara Pelepasan Alumni Fakultas Ilmu Komputer Universitas Sriwijaya ke-74', '2025-08-15 02:14:07.000', 'http://localhost:8080/assets/covers/1755198847385087700.jpg', '', 'admin web'),
(7, '2025-08-15 02:14:29.000', '2025-08-15 07:12:49.395', NULL, 'Secangkir Teh Bersama Ketua Jurusan Teknik Informatika', '2025-08-15 02:14:29.000', 'http://localhost:8080/assets/covers/1755198869166826400.jpg', '', 'admin web'),
(8, '2025-08-15 02:14:42.000', '2025-08-15 07:12:49.397', NULL, 'Pendaftaran Mahasiswa Baru S3 Doktor Fasilkom UNSRI 2024', '2025-08-15 02:14:42.000', 'http://localhost:8080/assets/covers/1755198882306015300.jpeg', '', 'admin web'),
(9, '2025-08-15 02:15:00.000', '2025-08-15 07:12:49.399', NULL, 'Universitas Sriwijaya (Unsri) melepas 16 orang mahasiswa asing yang berasal dari Universiti Teknikal Malaysia Melaka ', '2025-08-15 02:15:00.000', 'http://localhost:8080/assets/covers/1755198900320871700.jpg', '', 'admin web');

-- --------------------------------------------------------

--
-- Table structure for table `scholarships`
--

CREATE TABLE `scholarships` (
  `id` int NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `judul` longtext NOT NULL,
  `tanggal` date DEFAULT NULL,
  `cover` longtext,
  `content` longtext,
  `posted_by` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `scholarships`
--

INSERT INTO `scholarships` (`id`, `created_at`, `updated_at`, `deleted_at`, `judul`, `tanggal`, `cover`, `content`, `posted_by`) VALUES
(5, '2025-08-15 02:15:52.000', '2025-08-15 07:12:49.406', NULL, 'Taiwan Scholarship 2016', '2025-08-15', 'http://localhost:8080/assets/covers/1755198951801996900.png', '', 'admin web'),
(6, '2025-08-15 02:16:07.000', '2025-08-15 07:12:49.412', NULL, 'Program Beasiswa Satu Keluarga Satu Sarjana (SKSS)', '2025-08-15', 'http://localhost:8080/assets/covers/1755198966831035400.png', '', 'admin web'),
(7, '2025-08-15 02:16:26.000', '2025-08-15 07:12:49.415', NULL, 'PROGRAM BEASISWA PERTAMINA SOBAT BUMI 2016', '2025-08-15', 'http://localhost:8080/assets/covers/1755198986051613700.png', '', 'admin web'),
(8, '2025-08-15 02:16:40.000', '2025-08-15 07:12:49.417', NULL, 'PEMBUKAAN PENDAFTARAN AKADEMI FGA DIGITAL TALENT SCHOLARSHIP 2020', '2025-08-15', 'http://localhost:8080/assets/covers/1755199000028619600.png', '', 'admin web');

-- --------------------------------------------------------

--
-- Table structure for table `vacancies`
--

CREATE TABLE `vacancies` (
  `id` bigint UNSIGNED NOT NULL,
  `judul` longtext NOT NULL,
  `tanggal` datetime(3) DEFAULT NULL,
  `cover` longtext,
  `content` longtext,
  `posted_by` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `vacancies`
--

INSERT INTO `vacancies` (`id`, `judul`, `tanggal`, `cover`, `content`, `posted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
(5, 'PT. JASA RAHARJA membuka kesempatan bagi putra putri Indonesia untuk menjadi Pegawai PT. Jasa Raharja', '2025-08-15 02:17:07.000', 'http://localhost:8080/assets/covers/1755199026924965100.png', '', 'admin web', '2025-08-15 02:17:07.000', '2025-08-15 07:12:49.424', NULL),
(6, 'Recruitmen ASN Fakultas Ilmu Komputer Universitas Sriwijaya 2024', '2025-08-15 02:17:19.000', 'http://localhost:8080/assets/covers/1755199038712929900.png', '', 'admin web', '2025-08-15 02:17:19.000', '2025-08-15 07:12:49.428', NULL),
(7, 'CDC UNSRI berkolaborasi dengan TUGU INSURANCE', '2025-08-15 02:17:30.000', 'http://localhost:8080/assets/covers/1755199049693573300.png', '', 'admin web', '2025-08-15 02:17:30.000', '2025-08-15 07:12:49.430', NULL),
(8, 'Tes Wawancara Penerimaan Dosen Luar Biasa Fakultas Ilmu Komputer Tahun Ajaran Ganjil 2024-2025', '2025-08-15 02:17:47.000', 'http://localhost:8080/assets/covers/1755199067069856000.jpg', '', 'admin web', '2025-08-15 02:17:47.000', '2025-08-15 07:12:49.432', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_admins_username` (`username`),
  ADD KEY `idx_admins_deleted_at` (`deleted_at`);

--
-- Indexes for table `berita`
--
ALTER TABLE `berita`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_deleted_at` (`deleted_at`),
  ADD KEY `idx_berita_deleted_at` (`deleted_at`);

--
-- Indexes for table `files`
--
ALTER TABLE `files`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_files_deleted_at` (`deleted_at`);

--
-- Indexes for table `kemahasiswaan_kerjasamas`
--
ALTER TABLE `kemahasiswaan_kerjasamas`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_kemahasiswaan_kerjasamas_deleted_at` (`deleted_at`);

--
-- Indexes for table `scholarships`
--
ALTER TABLE `scholarships`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_scholarships_deleted_at` (`deleted_at`);

--
-- Indexes for table `vacancies`
--
ALTER TABLE `vacancies`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_vacancies_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `berita`
--
ALTER TABLE `berita`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `files`
--
ALTER TABLE `files`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `kemahasiswaan_kerjasamas`
--
ALTER TABLE `kemahasiswaan_kerjasamas`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `scholarships`
--
ALTER TABLE `scholarships`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `vacancies`
--
ALTER TABLE `vacancies`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
