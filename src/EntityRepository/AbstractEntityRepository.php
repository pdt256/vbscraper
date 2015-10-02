<?php
namespace pdt256\vbscraper\EntityRepository;

use Doctrine;
use Doctrine\ORM\QueryBuilder;
use pdt256\vbscraper\Entity\EntityInterface;

abstract class AbstractEntityRepository extends Doctrine\ORM\EntityRepository
{
    protected function createEntity(EntityInterface & $entity)
    {
        $entityManager = $this->getEntityManager();
        $entityManager->persist($entity);
        $entityManager->flush();
    }

    protected function updateEntity(EntityInterface & $entity)
    {
        $entityManager = $this->getEntityManager();
        $entity = $entityManager->merge($entity);
        $entityManager->flush();
    }

    protected function deleteEntity(EntityInterface $entity)
    {
        $entityManager = $this->getEntityManager();
        $entityManager->remove($entity);
        $entityManager->flush();
    }

    protected function persistEntity($entity)
    {
        $entityManager = $this->getEntityManager();
        $entityManager->persist($entity);
    }

    protected function flushEntity()
    {
        $entityManager = $this->getEntityManager();
        $entityManager->flush();
    }

    public function getQueryBuilder()
    {
        return new QueryBuilder($this->getEntityManager());
    }
}
