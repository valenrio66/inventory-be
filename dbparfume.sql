-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 10, 2023 at 05:28 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dbparfume`
--

-- --------------------------------------------------------

--
-- Table structure for table `parfume`
--

CREATE TABLE `parfume` (
  `id_parfume` bigint(20) NOT NULL,
  `nama_parfume` varchar(50) NOT NULL,
  `jenis_parfume` varchar(20) NOT NULL,
  `merk` varchar(30) NOT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `harga` int(10) NOT NULL,
  `thn_peluncuran` varchar(5) NOT NULL,
  `stok` int(5) NOT NULL,
  `ukuran` varchar(8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `parfume`
--

INSERT INTO `parfume` (`id_parfume`, `nama_parfume`, `jenis_parfume`, `merk`, `deskripsi`, `harga`, `thn_peluncuran`, `stok`, `ukuran`) VALUES
(1, 'Christian Dior', 'Sauvage', 'Christian Dior', 'Sauvage adalah parfum pria yang kuat dan maskulin. Parfum ini menampilkan aroma segar dan tajam dengan sentuhan citrus, lavender, dan kayu-kayuan. Dengan kombinasi yang unik, Sauvage memberikan kesan yang modern dan memikat.', 1000000, '1978', 200, '50 ml'),
(2, 'Christian Dior', 'Miss Dior', 'Christian Dior', 'Miss Dior adalah parfum ikonik dari Christian Dior yang diperkenalkan pada tahun 1947. Parfum ini memiliki aroma yang elegan, feminin, dan romantis. Biasanya dikemas dalam botol yang indah dengan pita berwarna merah sebagai tanda pengenalnya.', 1500000, '1947', 250, '50 ml'),
(3, 'Christian Dior', 'J\'adore', 'Christian Dior', 'J\'adore adalah parfum bunga yang mewah dan sensual. Aroma utamanya adalah bunga melati, mawar, dan ylang-ylang yang dikombinasikan dengan sentuhan buah-buahan dan kayu. Parfum ini mencerminkan keanggunan dan daya tarik yang klasik.', 2000000, '1947', 300, '50 ml'),
(4, 'Christian Dior', 'Hypnotic Poison', 'Christian Dior', 'Hypnotic Poison adalah parfum yang sensual dan misterius. Dengan campuran aroma vanila, mandarin, kemenyan, dan almond, parfum ini menciptakan kesan yang memikat dan memabukkan.', 2500000, '1878', 300, '50 ml');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id_user` bigint(20) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `username` varchar(10) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id_user`, `nama`, `username`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Valen Rionald', 'valen', '$2a$10$q55B5l/.n6xXiihZlGG/yed9D8btFzaoCwD4XRnvgvezvsTqArLvu', '2023-06-07 14:55:49', '2023-06-07 14:55:49', NULL),
(2, 'Richard', 'richard', '$2a$10$yk1BBBAMZTS/Q.eHVGUXhuhFJcoGL0aKDV4azZ2Zz1PvYhaORBzeu', '2023-06-07 14:56:58', '2023-06-07 14:56:58', NULL),
(3, 'testing', 'test', '$2a$10$Yu2b18Z4OkyA.Aib5HvEwuy7MHNJvnHPFbD3Lx7rpqeOW/DRXMXZq', '2023-06-08 11:02:45', '2023-06-08 11:02:45', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `parfume`
--
ALTER TABLE `parfume`
  ADD PRIMARY KEY (`id_parfume`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `parfume`
--
ALTER TABLE `parfume`
  MODIFY `id_parfume` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id_user` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
