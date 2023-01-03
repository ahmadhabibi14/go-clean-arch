CREATE TABLE mahasiswa (
   id INT(11) AUTO_INCREMENT NOT NULL,
   nama VARCHAR(255) NOT NULL,
   prodi VARCHAR(255) NOT NULL,
   semester INT(11) NOT NULL,
   PRIMARY KEY(id),
) ENGINE=InnoDB;

INSERT INTO `mahasiswa`(`nama`, `prodi`, `semester`)
VALUES (
   "Ahmad Rizky Nusantara Habibi",
   "Ilmu Komputer",
   2,
)