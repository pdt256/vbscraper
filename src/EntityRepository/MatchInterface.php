<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Match;

interface MatchInterface
{
    public function create(Match & $entity);
    public function update(Match & $entity);
    public function delete(Match & $entity);

    /**
     * @param int $id
     * @return Match
     */
    public function find($id);
}